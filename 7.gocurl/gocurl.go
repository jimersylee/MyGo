package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//初始化,在main函数之前会自动调用
func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./gocurl url")
		os.Exit(-1)
	}
}
//go的curl简单实现
func main() {
	response, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	//讲http请求返回的数据复制到标准输出
	io.Copy(os.Stdout, response.Body)
	err = response.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
}
