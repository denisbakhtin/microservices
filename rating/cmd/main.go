package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"
	"videoexample/gen"
	"videoexample/pkg/discovery"
	"videoexample/pkg/discovery/consul"
	"videoexample/rating/internal/controller"
	grpchandler "videoexample/rating/internal/handler/grpc"
	"videoexample/rating/internal/repository/mysql"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const serviceName = "rating"

func main() {
	var port int
	flag.IntVar(&port, "port", 8082, "API handler port")
	flag.Parse()
	log.Printf("Starting the video rating service on port %d", port)
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

	log.Println("Starting the rating service")
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
	gen.RegisterRatingServiceServer(srv, h)
	reflection.Register(srv)
	srv.Serve(lis)
}
