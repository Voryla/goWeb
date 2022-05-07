package nativeWeb

import (
	"fmt"
	"goWeb/entity"
	"goWeb/funcChain"
	"goWeb/useHttp2"
	"net/http"
)

func index(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello World!%s", request.URL.Path[1:])
}
func UseNativeWebHandler() {
	mux := http.NewServeMux()
	// 设置文件服务器
	// 目录结构需要注意
	// rootPath
	//	 ---public
	// 	 ---main.go
	files := http.FileServer(http.Dir("public/"))
	// 当服务器收到以/static/开头的URL请求时，会移除URL中的/static/字符串，然后在/public目录中查找被请求的文件
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", useHttp2.MyHandler{}.ServeHTTP)
	// 引入其他包中的HandleFunc
	//mux.HandleFunc("/", utils.HowDo)
	mux.HandleFunc("/user/add", funcChain.VerificationUser(entity.AddUser))
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
