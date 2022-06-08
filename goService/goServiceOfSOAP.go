package goService

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*
被解析的数据
<?xml version="1.0" encoding="utf-8"?>
<post id="1">
　<content>Hello World!</content>
　<author id="2">Sau Sheong</author>
</post>
*/

/*
解析规则
（1）通过创建一个名字为XMLName 、类型为xml.Name 的字段，可以将XML元素的名字存储在这个字段里面（在一般情况下，结构的名字就是元素的名字）。

（2）通过创建一个与XML元素属性同名的字段，并使用 'xml:"<name >,attr" '作为该字段的结构标签，可以将元素的<name > 属性的值存储到这个字段里面。

（3）通过创建一个与XML元素标签同名的字段，并使用 'xml:",chardata"' 作为该字段的结构标签，可以将XML元素的字符数据存储到这个字段里面。

（4）通过定义一个任意名字的字段，并使用'xml:",innerxml"' 作为该字段的结构标签，可以将XML元素中的原始XML存储到这个字段里面。

（5）没有模式标志（如,attr 、,chardata 或者,innerxml ）的结构字段将与同名的XML元素匹配。讲<>content</> 讲xml元素标签中的内容存储到这个字段中

（6）使用'xml:"a>b>c"' 这样的结构标签可以在不指定树状结构的情况下直接获取指定的XML元素，其中a 和b 为中间元素，而c 则是想要获取的节点元素。
	Comments []Comment `xml:"comments>comment"`
*/

// Post 用以解析对应xml的结构体
type Post struct {
	XMLName xml.Name `xml:"post"` // 规则1 解析XML元素的名字
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"` // 规则5
	Author  Author   `xml:"author"`
	Xml     string   `xml:",innerxml"` // 规则4 解析XML元素中的原始XML
}

// Author 用以解析对应xml的子结构体
type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"` // 讲xml元素的字符数据存储到该字段中：<author id="2">Sau Sheong</author> 元素中的<author id="2">Sau Sheong</author>
}

// ParseXML 通过xml.Unmarshal() 方法解析小型xml文件  XML文件较小使用
func ParseXML() {
	// 1. 打开文件
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()
	// 2. 读取数据
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML data:", err)
		return
	}
	var post Post
	// 3. 解析数据到结构体
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)
}

// ParseXMLWithStream 通过流高效解析XML  XML文件过大时使用
func ParseXMLWithStream() {
	xmlFile, err := os.Open("postOutput.xml")
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()
	// 通过文件获取decoder
	decoder := xml.NewDecoder(xmlFile)
	var post Post
	for {
		// 获取一个xml节点
		t, err := decoder.Token()
		// 当文件流读取完
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML into tokens:", err)
			return
		}
		switch se := t.(type) {
		// 判断该token是否为XML元素的起始标签
		case xml.StartElement:
			// 节点的元素名称
			if se.Name.Local == "post" {
				// 开始解析
				decoder.DecodeElement(&post, &se)
			}
		}
	}
	fmt.Println(post)
}

// CreateXML 创建XML文件
func CreateXML() {
	// 1. 创建数据结构体
	post := Post{
		Id:      "1",
		Content: "Hello World!",
		Author: Author{
			Id:   "2",
			Name: "Joe",
		},
	}
	// 2. 编码结构体
	output, err := xml.Marshal(&post)
	if err != nil {
		fmt.Println("Error marshalling to XML:", err)
		return
	}
	// 3. 写入文件
	//err = ioutil.WriteFile("postOutput.xml", output, 0644)
	err = ioutil.WriteFile("postOutput.xml", []byte(xml.Header+string(output)), 0644)
	if err != nil {
		fmt.Println("Error writing XML to file:", err)
		return
	}
}

func CreateXMLWithStream() {
	// 1. 创建数据结构体
	post := Post{
		Id:      "1",
		Content: "Hello World!",
		Author: Author{
			Id:   "2",
			Name: "我",
		},
	}

	stream, _ := os.Open("streamPostOutput.xml")
	stream.Write([]byte(xml.Header))
	//stream, _ := os.Create("streamPostOutput.xml")
	defer stream.Close()
	encoder := xml.NewEncoder(stream)
	encoder.Indent("", "\t")
	encoder.Encode(&post)
	encoder.Flush()

}
