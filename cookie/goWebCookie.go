package cookie

import (
	"encoding/base64"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func RunWebOfLearnCookie() {
	router := httprouter.New()
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}
	router.GET("/setCookie", setCookie)
	router.GET("/getCookie", getCookie)
	router.GET("/delCookie", delCookie)
	server.ListenAndServe()
}

// setCookie 设置cookie
func setCookie(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming%20",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    base64.URLEncoding.EncodeToString([]byte("Manning Publications Co %B3")),
		HttpOnly: true,
	}
	c3 := http.Cookie{
		Name:     "third_cookie",
		Value:    "Manning Publications Co",
		HttpOnly: true,
	}
	header := w.Header()
	header.Set("Set-Cookie", c1.String())
	header.Add("Set-Cookie", c2.String())
	// Go语言提供的更便捷设置cookie的方法
	http.SetCookie(w, &c3)
}

// getCookie 获取cookie
func getCookie(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//返回了一个切片，这个切片包含了一个字符串，而这个字符串又包含了客户端发送的任意多个cookie
	cookies := r.Header["Cookie"]
	fmt.Fprintln(w, cookies)
	// 便捷的获取指定key的cookie值
	cookie, _ := r.Cookie("first_cookie")
	fmt.Fprintln(w, cookie)
	secondCookie, _ := r.Cookie("second_cookie")
	secondCookieStr, _ := base64.URLEncoding.DecodeString(secondCookie.Value)
	fmt.Fprintln(w, string(secondCookieStr))
	// 获取所有cookie，返回值为*http.Cookie的切片
	//r.Cookies()
	// cookie 添加到请求中
	//r.AddCookie(&http.Cookie{})
}

// delCookie 删除cookie
func delCookie(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	cookie, err := r.Cookie("first_cookie")
	if err == nil {
		// 删除指定的cookie
		cookie.MaxAge = -1
		cookie.Expires = time.Unix(1, 0)
		// 替换客户端本地的该cookie.Name的cookie，即将最大过期时间MaxAge和Expires设置为一个已经过去的时间，
		// 当浏览器接收该cookie时检测其过期时间就会删除该Name的Cookie
		http.SetCookie(w, cookie)
	}
}
