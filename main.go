package main

import (
	"log"
	"net/http"
)

func main() {
	// 設置靜態文件目錄為根目錄
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	// 啟動服務器
	log.Printf("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
