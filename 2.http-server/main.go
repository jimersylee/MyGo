package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

/**
http包建立web服务器
main入口函数,只匹配根路径
 */
/*func main() {
	http.HandleFunc("/", sayHelloName)       //设置访问路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}*/

/**
处理路由,输出hello world
 */
func sayHelloName(responseWriter http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Println("param:", request.Form)    //调试输出请求参数信息
	fmt.Println("path:", request.URL.Path) //调试输出请求url信息

	for k, v := range request.Form {
		fmt.Println("key:", k)
		fmt.Println("value", strings.Join(v, "")) //将value处理成string
	}

	fmt.Fprint(responseWriter, "Hello, Jimersy Lee!")

}

/**
实现简单的路由
 */



/**
定义一个指针结构体指向自定义的ServeHTTP
 */
type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//访问路径为根目录
	if r.URL.Path == "/" {
		sayHelloName(w, r)
		return
	}

	//其他的路由

	//如果没有匹配到
	http.NotFound(w, r)
	return
}




func main() {
	mux := &MyMux{}
	err:=http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

