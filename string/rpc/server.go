package main

import (
	"flag"
	"github.com/FFFcomewhere/University-course-selection/string/rpc/pb"
	string_service "github.com/FFFcomewhere/University-course-selection/string/rpc/string-service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalln("falied to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	strintService := new(string_service.StringService)
	pb.RegisterStringServiceServer(grpcServer, strintService)
	grpcServer.Serve(lis)

}
