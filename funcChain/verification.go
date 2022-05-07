package funcChain

import (
	"fmt"
	"net/http"
	"strconv"
)

func VerificationUser(proc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		method := request.Method
		fmt.Println(method)
		name := request.FormValue("name")
		age := request.FormValue("age")
		_, err := strconv.ParseInt(age, 10, 8)
		if name == "" || err != nil {
			writer.WriteHeader(404)
			return
		}
		proc(writer, request)
	}
}
