package goTesting

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
)

// PostJson 用以存储Json数据的结构
type PostJson struct {
	Id       int           `json:"id"`
	Content  string        `json:"content"`
	Author   AuthorJson    `json:"author"`
	Comments []CommentJson `json:"comments"`
}

type AuthorJson struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CommentJson struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func ParseJsonWithDecoder(filename string) (post PostJson, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&post)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	return
}

func ParseJsonWithUnmarshal() {
	// 打开文件
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	// 关闭文件
	defer jsonFile.Close()
	// 读取数据
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}
	post := PostJson{}
	json.Unmarshal(jsonData, &post)
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
	case "PUT":
	case "DELETE":
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	if id == 0 {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	post, _ := ParseJsonWithDecoder("post.json")
	jsonData, _ := json.MarshalIndent(post, "", "\t")
	w.Write(jsonData)
	return
}
