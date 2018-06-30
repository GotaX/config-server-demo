package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/search", func(w http.ResponseWriter, req *http.Request) {
		q := req.URL.Query().Get("q")
		src := fmt.Sprintf("https://cn.bing.com/search?q=%v", q)
		fmt.Fprintf(w, `<iframe width="100%%" height="98%%" scrolling="auto" frameborder="0" src="%v">`, src)
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
