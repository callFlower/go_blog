package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	// 如果这里是title，包外代码将无法访问或修改这个字段，尽管可以访问结构体本身
	Title string `json:"title"` // Title代表外部可访问 title代表最终响应的内容
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, req *http.Request) {
	// 如果不做设置 -- Content-Type: text/plain; charset=utf-8
	// 考虑到前后端大部分以json格式传输 -- Content-Type: application/json
	w.Header().Set("Content-Type", "application/json")
	var indexdata IndexData
	indexdata.Title = "go_blog"
	indexdata.Desc = "learning coding"
	jsonstr, err := json.Marshal(indexdata)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}
	w.Write(jsonstr)
	// w.Write([]byte("hello go blog"))
}

func indexHtml(w http.ResponseWriter, req *http.Request) {
	var indexdata IndexData
	indexdata.Title = "go_blog"
	indexdata.Desc = "learning coding"
	// 名字要与模版名一致才行
	t := template.New("index.html")

	path, _ := os.Getwd()
	t, _ = t.ParseFiles(path + "/template/index.html")
	t.Execute(w, indexdata)
	// w.Write([]byte("hello go blog"))
}

func main() {
	// http - ip port
	server := http.Server{
		// 只赋值部分字段时，其他未被显式赋值的字段将会被初始化为其零值。
		Addr: "127.0.0.1:8080",
	}

	// 输入127.0.0.1:8080 表示 127.0.0.1:8080/ 也就是访问的是根目录
	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", indexHtml)

	// ListenAndServe: 监听端口+启动服务
	// 如果出错，打印报错
	if err := server.ListenAndServe(); err != nil {
		log.Print(err)
	}
}
