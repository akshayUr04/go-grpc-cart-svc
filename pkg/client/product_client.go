package client

import (
	"context"
	"fmt"

	"github.com/akshayUr04/go-grpc-cart-svc/pkg/pb"
	"google.golang.org/grpc"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func InintProductServiceClient(url string) ProductServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := ProductServiceClient{
		Client: pb.NewProductServiceClient(cc),
	}
	return c
}

func (c *ProductServiceClient) FindOne(productId int64) (*pb.FindOneResponse, error) {
	req := &pb.FindOneRequest{
		Id: productId,
	}
	return c.Client.FindOne(context.Background(), req)
}
