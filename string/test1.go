package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("hello")
	})
	fmt.Println(" i am test1")
	err := http.ListenAndServe(":8080", handler)
	log.Fatalln(err)

}
