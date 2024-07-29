package main

import (
	"log"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	// 如果不做设置 -- Content-Type: text/plain; charset=utf-8
	// 考虑到前后端大部分以json格式传输 -- Content-Type: application/json
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("hello go blog"))
}

func main() {
	// http - ip port
	server := http.Server{
		// 只赋值部分字段时，其他未被显式赋值的字段将会被初始化为其零值。
		Addr: "127.0.0.1:8080",
	}

	// 输入127.0.0.1:8080 表示 127.0.0.1:8080/ 也就是访问的是根目录
	http.HandleFunc("/", index)

	// ListenAndServe: 监听端口+启动服务
	// 如果出错，打印报错
	if err := server.ListenAndServe(); err != nil {
		log.Print(err)
	}
}
