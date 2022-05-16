package goFileSystem

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type post struct {
	Id      int
	Content string
	Author  string
}

func goUseCsv() {
	// 创建file
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	allPosts := []post{
		{Id: 1, Content: "Hello World1", Author: "Sau She1"},
		{Id: 2, Content: "Hello World2", Author: "Sau She2"},
		{Id: 3, Content: "Hello World3", Author: "Sau She3"},
		{Id: 4, Content: "Hello World4", Author: "Sau She4"},
	}
	// 通过 file 创建 writer
	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		// 写入一行数据，数据为string切片
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	// 写入磁盘
	writer.Flush()

	// 打开指定文件 file
	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 通过 file 创建 reader
	reader := csv.NewReader(file)
	// FieldsPerRecord 、
	// 为负：不检查每行字段数
	// 为0：将读取的第一行记录中拥有的字段数作为后续预期值
	// 为正：将该值作为每行记录预期字段数
	reader.FieldsPerRecord = 3
	// 读取所有数据返回 二维切片
	record, err := reader.ReadAll()
	// 在读取完数据后，FieldPos 可以返回给定第n(从0开始)个字段的行号和列号
	//line, colm := reader.FieldPos(1)
	//fmt.Println(line, colm)
	//reader.Read() // 单行读取
	if err != nil {
		panic(err)
	}
	var posts []post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := post{
			Id:      int(id),
			Content: item[1],
			Author:  item[2],
		}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
