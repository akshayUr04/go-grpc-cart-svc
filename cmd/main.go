package main

import (
	"fmt"
	"log"
	"net"

	"github.com/akshayUr04/go-grpc-cart-svc/pkg/client"
	"github.com/akshayUr04/go-grpc-cart-svc/pkg/config"
	"github.com/akshayUr04/go-grpc-cart-svc/pkg/db"
	"github.com/akshayUr04/go-grpc-cart-svc/pkg/pb"
	"github.com/akshayUr04/go-grpc-cart-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("failed to listing", err)
	}

	productSvc := client.InintProductServiceClient(c.ProductSvcUrl)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("cart Svc on", c.Port)

	s := services.Server{
		H:          h,
		ProductSvc: productSvc,
	}
	grpcServer := grpc.NewServer()

	pb.RegisterCartServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
