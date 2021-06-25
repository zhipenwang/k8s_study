package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("HTTPPORT")
	name, _ := os.Hostname()
	http.HandleFunc("/demo", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello k8s-golang, podName:" + name))
	})
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello k8s-golang-test, podName:" + name))
	})
	http.ListenAndServe(":"+port, nil)
}
