package main

import (
	"goWeb/goTemplate"
)

func main() {
	// 原生http服务器
	//nativeWeb.UseNativeWebHandler()
	// 使用httpRouter框架的http服务器
	//useHttpRuouter.NewRouter()
	// 使用http2
	//useHttp2.UseHttp2()

	// 查看 Request 结构中的URL结构体内容
	//goWebRequestAndResponse.RunWeb()

	// cookie
	//cookie.RunWebOfLearnCookie()

	// template
	goTemplate.GoWebRunTemplate()
}
