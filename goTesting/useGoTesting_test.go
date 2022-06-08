package goTesting // 测试文件与被测试的源代码文件位于同一个包内

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

// 功能性测试

func TestDecode(t *testing.T) {
	post, err := ParseJsonWithDecoder("post.json") // 调用被测试函数
	if err != nil {
		t.Error(err)
	}
	// 检查结果是否和预期的一样，如果不一样就显示一条错误
	if post.Id != 1 {
		t.Error("Wrong id, was expecting 1 but got", post.Id)
	}
	if post.Content != "Hello World!" {
		t.Error("Wrong content, was expecting 'Hello World!' but got", post.Content)
	}
}

func TestEncode(t *testing.T) {
	// 暂时跳过对此函数的测试
	t.Skip("Skipping encoding for now")
	t.Fatal()
}

// testShort 根据-short条件决定是否跳过测试用例
func TestShort(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test in short mode")
	}
	time.Sleep(10 * time.Second)
}

// 并行测试

func TestParallel_1(t *testing.T) { // 模拟耗时1秒钟的任务
	t.Parallel() // 调用Parallel函数，以并行方式运行测试用例
	time.Sleep(1 * time.Second)
}

func TestParallel_2(t *testing.T) { // 模拟耗时2秒钟的任务
	t.Parallel()
	time.Sleep(2 * time.Second)
}

func TestParallel_3(t *testing.T) { // 模拟耗时3秒钟的任务
	t.Parallel()
	time.Sleep(3 * time.Second)
}

// 基准测试,对比以下两个基准测试，可以看出对Json的解析，Decode 将快于 Unmarshal

func BenchmarkWithDecoder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseJsonWithDecoder("post.json")
	}
}

func BenchmarkParseJsonWithUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseJsonWithUnmarshal()
	}
}

// go web 测试

func TestHandleGet(t *testing.T) {
	mux := http.NewServeMux()                            // 创建一个用于运行测试的多路复用器
	mux.HandleFunc("/post/", HandleRequest)              // 绑定想要测试的处理器
	writer := httptest.NewRecorder()                     // 创建记录器，用于获取服务器返回的HTTP响应
	request, _ := http.NewRequest("GET", "/post/1", nil) // 为被测试的处理器创建相应的请求
	mux.ServeHTTP(writer, request)                       // 向被测试的处理器发送请求
	if writer.Code != 200 {                              // 对记录器记载的响应结果进行检查
		t.Errorf("Response code is %v", writer.Code)
	}
	var post PostJson
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Error("Cannot retrieve JSON post")
	}
	fmt.Println(post)
}

// 为了保持代码的简洁性，我们可以把一些重复出现的测试代码以及其他测试夹具（fixture）设置在TestMain函数中，Go的testing包允许用户通过TestMain 函数在进行测试时执行相应的预设(setUp)操作或者拆卸(teardown)操作。
// setUp 函数和tearDown 函数是为所有测试用例设置的，
// 它们在整个测试过程中只会被执行一次，其中setUp 函数会在所有测试用例被执行之前执行，而tearDown 函数则会在所有测试用例都被执行完毕之后执行。
func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}
func setUp() {
	fmt.Println("up")
}

func tearDown() {
	fmt.Println("down")
}

func TestName(t *testing.T) {

}
