package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("HTTPPORT")
	http.HandleFunc("/demo", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello k8s-golang"))
	})
	http.ListenAndServe(":"+port, nil)
}
