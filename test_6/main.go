package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// 打开您输入的类型的文件
	// 想检查
	file, err := os.Open("D://de518e7a1d02eab0ab20818e01e3c623.png")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	// 获取文件内容
	contentType, err := GetFileContentType(file)

	if err != nil {
		panic(err)
	}

	fmt.Println("文件的内容类型是： " + contentType)
}

func GetFileContentType(ouput *os.File) (string, error) {

	// 仅嗅探内容类型的第一个
	// 使用了 512 个字节。

	buf := make([]byte, 512)

	_, err := ouput.Read(buf)

	if err != nil {
		return "", err
	}


	// 真正起作用的函数
	contentType := http.DetectContentType(buf)

	return contentType, nil
}