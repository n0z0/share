package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println(SetPort())
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.ListenAndServe(":"+SetPort(), nil)
}

func SetPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "1024"
	}
	return port
}
