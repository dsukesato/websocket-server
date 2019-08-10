package main

import (
	"log"
	"net/http"

	"github.com/dsukesato/websocket-server/models"
)


func main() {
	// localhost:8080 でアクセスした時に index.html を読み込む
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/chat", models.HandleClients)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("error starting http server::", err)
		return
	}
}


