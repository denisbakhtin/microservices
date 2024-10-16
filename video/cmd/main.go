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
	"videoexample/video/internal/controller"

	metadatagateway "videoexample/video/internal/gateway/metadata/grpc"
	ratinggateway "videoexample/video/internal/gateway/rating/grpc"
	grpchandler "videoexample/video/internal/handler/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const serviceName = "video"

func main() {
	var port int
	flag.IntVar(&port, "port", 8083, "API handler port")
	flag.Parse()
	log.Printf("Starting the video service on port %d", port)
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
	metadataGateway := metadatagateway.New(registry)
	ratingGateway := ratinggateway.New(registry)
	ctrl := controller.New(ratingGateway, metadataGateway)
	h := grpchandler.New(ctrl)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	reflection.Register(srv)
	gen.RegisterVideoServiceServer(srv, h)
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}
