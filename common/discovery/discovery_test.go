package discovery

import (
	"log"
	"testing"
	"time"
)

func TestDiscovery(t *testing.T) {
	var endpoints = []string{"localhost:2379"}
	service := NewServiceDiscovery(endpoints)
	defer service.Close()
	service.WatchService("/web/")
	service.WatchService("/gRPC")
	service.WatchService("/time")
	for {
		select {
		case <-time.Tick(10 * time.Second):
			log.Println(service.GetSService())
		}
	}
}
