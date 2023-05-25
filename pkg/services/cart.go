package services

import (
	"github.com/akshayUr04/go-grpc-cart-svc/pkg/client"
	"github.com/akshayUr04/go-grpc-cart-svc/pkg/db"
	"github.com/akshayUr04/go-grpc-cart-svc/pkg/pb"
)

type Server struct {
	H          db.Handler
	ProductSvc client.ProductServiceClient
	pb.UnimplementedCartServiceServer
}
