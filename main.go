package main

import (
	"fmt"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello, 欢迎来到 goblog！</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>")
	}
}

func aboutHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "<h1>关于我们</h1>")
}

func main() {
	// 重构 HandleFunc 的底层 ServeMux 的实现
	router := http.NewServeMux()

	router.HandleFunc("/", defaultHandler)
	router.HandleFunc("/about", aboutHandler)

	// http.HandleFunc("/", defaultHandler)
	// http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8080", router)
}
