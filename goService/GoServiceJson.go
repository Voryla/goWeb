package goService

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
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

// ParseJson using json.Unmarshal
func ParseJson() {
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
	fmt.Println(post)
}

// ParseJsonWithDecoder using json.NewDecoder
func ParseJsonWithDecoder() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error opening Json file:", err)
		return
	}
	defer jsonFile.Close()
	// 根据文件创建相应的解码器
	decoder := json.NewDecoder(jsonFile)
	// 遍历JSON文件，直到遇见EOF为止
	for {
		var post PostJson
		// 将数据解吗至结构体
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}
		fmt.Println(post)
	}
}

// CreateJson 使用json.MarshalIndent编码json数据，将Json写入文件
func CreateJson() {
	post := PostJson{
		Id:      1,
		Content: "Hello World!",
		Author: AuthorJson{
			Id:   2,
			Name: "Sau Sheong",
		},
		Comments: []CommentJson{
			{
				Id:      3,
				Content: "Have a great day!",
				Author:  "Adam",
			},
			{
				Id:      4,
				Content: "How are you today?",
				Author:  "Betty",
			},
		},
	}
	// 设置缩进并编码数据
	outputData, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling to Json:", err)
		return
	}
	// 将数据写入文件
	ioutil.WriteFile("postOutput.json", outputData, 0644)
	if err != nil {
		fmt.Println("Error writing Json to file:", err)
		return
	}
}

// CreateJsonWithEncoder 使用Encoder编码Json数据，并保存至文件中
func CreateJsonWithEncoder() {
	post := PostJson{
		Id:      1,
		Content: "Hello World!",
		Author: AuthorJson{
			Id:   2,
			Name: "Sau Sheong",
		},
		Comments: []CommentJson{
			{
				Id:      3,
				Content: "Have a great day!",
				Author:  "Adam",
			},
			{
				Id:      4,
				Content: "How are you today?",
				Author:  "Betty",
			},
		},
	}
	jsonFile, err := os.Create("postEncoderOutput.json")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return
	}
}
