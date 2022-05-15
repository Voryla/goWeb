package goTemplate

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

func GoWebRunTemplate() {
	router := httprouter.New()
	router.GET("/helloTemplate", helloTemplate)
	router.GET("/stringTemplate", stringTemplate)
	router.GET("/chooseTemplate", chooseTemplate)
	router.GET("/templateIf", templateIf)
	router.GET("/templateFor", templateFor)
	router.GET("/templateSet", templateSet)
	router.GET("/templateInclude", templateInclude)
	router.GET("/templateWithFunc", templateWithFunc)
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

// stringTemplate 解析String模板
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

// chooseTemplate 解析多个模板,模板引擎可以解析多个模板文件，但是可以选择其中一个使用
func chooseTemplate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	t, _ := template.ParseFiles("template/helloTemplate.html", "template/helloTemplate2.html")
	t.ExecuteTemplate(w, "helloTemplate2.html", "chooseTemplate")
}

// templateIf 模板引擎条件动作
func templateIf(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	rand.Seed(time.Now().Unix())
	t, _ := template.ParseFiles("template/templateIf.html")
	t.Execute(w, rand.Intn(10) > 5)
}

// templateFor 模板引擎迭代动作
func templateFor(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	t, _ := template.ParseFiles("template/templateFor.html")
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	t.Execute(w, daysOfWeek)
}

// templateSet 模板引擎设置动作
func templateSet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	t, _ := template.ParseFiles("template/templateSet.html")
	t.Execute(w, "hello")
}

// templateInclude 模板引擎包含动作
func templateInclude(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	t, _ := template.ParseFiles("template/includeTemplate1.html", "template/includeTemplate2.html")
	t.Execute(w, "hello")
}
func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

// templateWithFunc 模板引擎调用函数
func templateWithFunc(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("templateWithFunc.html").Funcs(funcMap)
	t.ParseFiles("template/templateWithFunc.html")
	t.Execute(w, time.Now())
}
