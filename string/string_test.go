package main

import (
	"fmt"
	"github.com/FFFcomewhere/University-course-selection/common/discovery"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestBase(t *testing.T) {

	handler := http.NewServeMux()
	handler.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("hello")
	})
	fmt.Println(" i am test1")
	err := http.ListenAndServe(":8080", handler)
	log.Fatalln(err)
}

func TestDiscovery(t *testing.T) {
	var endpoints = []string{"localhost:2379"}
	service := discovery.NewServiceDiscovery(endpoints)
	defer service.Close()
	service.WatchService("/web/")
	service.WatchService("/gRPC")
	service.WatchService("/string1")
	for {
		select {
		case <-time.Tick(10 * time.Second):
			log.Println(service.GetSService())
		}
	}
}
