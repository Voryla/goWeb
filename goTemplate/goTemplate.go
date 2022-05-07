package goTemplate

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"text/template"
)

func GoWebRunTemplate() {
	router := httprouter.New()
	router.GET("/helloTemplate", helloTemplate)
	router.GET("/stringTemplate", stringTemplate)
	server := http.Server{Addr: "0.0.0.0:8080", Handler: router}
	server.ListenAndServe()
}

// helloTemplate 展示模板
func helloTemplate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// 对模板文件进行语法分析
	t, err := template.ParseFiles("template/helloTemplate.html")
	if err == nil {
		// 将 ResponseWriter 和数据一起传入Execute方法,即模板引擎在生成HTML后就可以把该HTML传给 ResponseWriter 了
		t.Execute(w, "Hello World!")
	} else {
		fmt.Fprintln(w, err.Error())
	}
}

func stringTemplate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	temp := `<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>Title</title>
			</head>
			<body>
			  {{ . }}
			</body>
			</html>`
	te := template.New("template")
	te.Parse(temp)
	te.Execute(w, "hello world")
}
