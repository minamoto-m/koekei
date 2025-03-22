package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "トップ画面へようこそ！")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("サーバーがポート8080で起動しました...")
	http.ListenAndServe(":8080", nil)
}
