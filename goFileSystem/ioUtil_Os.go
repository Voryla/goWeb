package goFileSystem

import (
	"fmt"
	"io/ioutil"
	"os"
)

func useIoUtiltil() {
	// 将要写入的数据作为切片传入
	err := ioutil.WriteFile("data1", []byte("Hello World!\n"), 0644)
	if err != nil {
		panic(err)
	}

	// 将读取的数据作为切片返回
	read1, _ := ioutil.ReadFile("data1")
	fmt.Print(string(read1))
}

func useOs() {
	data := []byte("Hello World! \n")
	file1, _ := os.Create("data2")
	defer file1.Close()

	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	file2, _ := os.Open("data2")
	defer file2.Close()

	read2 := make([]byte, len(data)+2)
	readSize, _ := file2.Read(read2)
	read2 = read2[:readSize]

	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(read2))
}
