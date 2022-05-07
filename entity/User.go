package entity

import (
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	name string
	age  uint8
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	age, _ := strconv.ParseInt(r.FormValue("age"), 10, 0)
	u := User{
		name: name,
		age:  uint8(age),
	}
	fmt.Fprintln(w, u)
}
