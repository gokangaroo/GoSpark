package main

import (
	"net/http"
	"log"
	"fmt"
	)

func home(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()  // 解析参数，默认是不会解析的
	fmt.Fprintf(w, "Hello GoSpark!") // 这个写入到 w 的是输出到客户端的
}
func main()  {
	http.HandleFunc("/",home)// 设置访问的路由
	err := http.ListenAndServe(":9000",nil)// 设置监听的端口
	if err != nil{
		log.Fatal("ListenAndServer: ",err)
	}
}

