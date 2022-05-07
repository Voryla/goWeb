package useHttpRuouter

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func NewRouter() {
	// 创建一个多路复用器
	mux := httprouter.New()
	// 将处理器函数与给定的HTTP方法进行绑定,  :name 为具名参数
	mux.GET("/hello/:name", helloHttpRouter)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

// get 方法 第三个参数为具名参数slice 即/hello/:name中 :name 这种形式的URL
func helloHttpRouter(response http.ResponseWriter, request *http.Request, param httprouter.Params) {
	fmt.Fprintf(response, "Hello World!%s", param.ByName("name"))
}
