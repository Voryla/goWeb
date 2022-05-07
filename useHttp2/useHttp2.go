package useHttp2

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
func UseHttp2() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	// 路径相对于调用者，该函数的调用者为main，所以路径要基于main.go与文件的相对路径
	server.ListenAndServeTLS("utils/cert.pem", "utils/key.pem")
}
