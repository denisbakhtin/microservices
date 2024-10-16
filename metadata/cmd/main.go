package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"
	"videoexample/gen"
	"videoexample/metadata/internal/controller"
	grpchandler "videoexample/metadata/internal/handler/grpc"
	"videoexample/metadata/internal/repository/mysql"
	"videoexample/pkg/discovery"
	"videoexample/pkg/discovery/consul"

	"google.golang.org/grpc"
)

const serviceName = "metadata"

func main() {
	var port int
	flag.IntVar(&port, "port", 8081, "API handler port")
	flag.Parse()
	log.Printf("Starting the video metadata service on port %d", port)
	registry, err := consul.NewRegistry("172.17.0.1:8500")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		panic(err)
	}
	go func() {
		for {
			if err := registry.ReportHealthyState(instanceID, serviceName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer registry.Deregister(ctx, instanceID, serviceName)
	repo, err := mysql.New()
	if err != nil {
		panic(err)
	}
	ctrl := controller.New(repo)
	h := grpchandler.New(ctrl)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	gen.RegisterMetadataServiceServer(srv, h)
	srv.Serve(lis)
}
