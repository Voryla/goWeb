package goWebRequestAndResponse

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

type user struct {
	Username string
	Pwd      string
}

func RunWeb() {
	router := httprouter.New()
	router.GET("/printUrl/* ", PrintRequestURL)
	router.GET("/printHeader/* ", PrintRequestHeader)
	router.POST("/printBody/* ", PrintRequestBody)
	router.POST("/printForm/* ", PrintRequestForm)
	router.POST("/printPostForm/* ", PrintRequestPostForm)
	router.POST("/printMultipartForm/* ", PrintRequestMultipartForm)
	router.POST("/uploadFile", uploadFile)
	router.GET("/aboutResponseWriter", aboutResponseWriter)
	router.GET("/jsonExample", jsonExample)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}
	server.ListenAndServe()
}

// PrintRequestURL 获取请求报文中的 URL 请求行
func PrintRequestURL(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "Scheme:%s\n", r.URL.Scheme)
	fmt.Fprintf(w, "Opaque:%s\n", r.URL.Opaque)
	fmt.Fprintf(w, "User:%s\n", r.URL.User)
	fmt.Fprintf(w, "Host:%s\n", r.URL.Host)
	fmt.Fprintf(w, "Path:%s\n", r.URL.Path)
	fmt.Fprintf(w, "RawPath:%s\n", r.URL.RawPath)
	fmt.Fprintf(w, "ForceQuery:%v\n", r.URL.ForceQuery)
	fmt.Fprintf(w, "RawQuery:%s\n", r.URL.RawQuery)
	fmt.Fprintf(w, "Fragment:%s\n", r.URL.Fragment)
	fmt.Fprintf(w, "RawFragment:%s\n", r.URL.RawFragment)
}

// PrintRequestHeader 获取请求报文中的 Header 请求头
func PrintRequestHeader(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	// 获取Header，Header是一个map map[string][]string
	header := r.Header
	// 通过原生map方式获取指定key的value
	accpetEncodingOfSlice := header["Accept-Encoding"]
	// 使用Header提供的Get方法获取指定key中的Slice的第一个值
	accpetEncodingOfString := header.Get("Accept-Encoding")
	fmt.Fprintln(w, accpetEncodingOfSlice)
	fmt.Fprintln(w, accpetEncodingOfString)
	// 将值附加到指定key指向的slice中
	header.Add("Allow", "DELETE")
	header.Add("Allow", "PUT")
	fmt.Fprintln(w, header.Get("Allow"))
	fmt.Fprintln(w, header["Allow"])
	// 设置并替换Header中指定key中的值
	header.Set("Allow", "POST")
	fmt.Fprintln(w, header.Get("Allow"))
	// 删除Header中指定key-value
	header.Del("Allow")
	fmt.Fprintln(w, header.Get("Allow"))
}

// PrintRequestBody 获取请求报文中的 Body 请求体
func PrintRequestBody(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
	user1 := user{}
	json.Unmarshal(body, &user1)
	fmt.Fprintln(w, user1.Username)
	fmt.Fprintln(w, user1.Pwd)
}

// PrintRequestForm 获取请求报文中的查询字段和请求体
func PrintRequestForm(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	r.ParseForm()
	// 同时返回URL参数和请求体的内容
	fmt.Fprintln(w, r.Form)
	// 如果查询字段与请求体同时包含某个字段，在使用Form查询值时，两者的值会放到同一个key中，且查询字段值会在请求体值之后
	// 如下，URL查询字段 hello=world ，请求体字段 hello = sau 则
	fmt.Fprintln(w, r.Form["hello"][0]) // output: sau
	fmt.Fprintln(w, r.Form["hello"][1]) // output: world
}

// PrintRequestPostForm 仅获取请求报文中的form表单内容
func PrintRequestPostForm(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	r.ParseForm()
	// 尝试请求：http://localhost:8080/printPostForm/?username=3
	// 同时请求体为application/x-www-form-urlencoded 格式，存在一个参数 hello = 3， 则使用PostForm仅会获取请求体中的内容，而不会获取查询字段username的内容
	fmt.Fprintf(w, "PostForm:%v\n", r.PostForm)
	fmt.Fprintf(w, "Form:%v\n", r.Form)
}

// PrintRequestMultipartForm 由于PostForm、Form字段仅支持application/x-www-form-urlencoded编码格式，
// 当Form使用multipart/form-data 编码时则需要使用ParseMultipartForm 方法和MultipartForm 字段
func PrintRequestMultipartForm(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "Form:%v\n", r.FormValue("hello"))
	fmt.Fprintf(w, "PostForm:%v\n", r.PostFormValue("hello"))
	fmt.Fprintf(w, "MultipartForm:%v\n", r.MultipartForm)
}

// uploadFile 使用MultipartForm字段接收上传的文件
func uploadFile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// 解析表单 同样可以使用 r.FormFile()获取FormFile()方法中调用了ParseMultipartForm()方法
	r.ParseMultipartForm(1024)
	// 取出name 为 uploaded 的第0个文件头
	fileHeader := r.MultipartForm.File["uploaded"][0]
	// 打开文件
	file, err := fileHeader.Open()
	if err == nil {
		// 将文件内容读取到字节数组中
		data, err := ioutil.ReadAll(file)
		if err == nil {
			// 这里模拟打印，实际操作中可能涉及write保存操作
			fmt.Fprintln(w, string(data))
		}
	}
}

// aboutResponseWriter 操作http报文
// http.ResponseWriter 是一个接口，其真正调用的实现着为 http.server.response
func aboutResponseWriter(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	str := `<html>
	<head><title>GoWebProgramming</title></head>
	<body><h1>Hello World</h1></body>
	</html>`
	header := w.Header()
	header.Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(200)         // 设置响应状态码，设置之后将无法改变响应头
	header.Set("Timeout", "3") // 此行对响应头的设置将不会生效，因为WriteHeader已经被调用
	// 通过设置响应头和响应码的配合可以实现重定向
	//w.Header().Set("Location", "http://google.com")
	//w.WriteHeader(302)
	w.Write([]byte(str)) // 设置响应状态码不影响后续对响应体的写入操作

}

type Post struct {
	User    string
	Threads []string
}

// jsonExample 通过设置响应头返回json数据
func jsonExample(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Sau Sheong",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}
