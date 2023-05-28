package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/akshayUr04/go-grpc-cart-svc/pkg/client"
	"github.com/akshayUr04/go-grpc-cart-svc/pkg/db"
	"github.com/akshayUr04/go-grpc-cart-svc/pkg/models"
	"github.com/akshayUr04/go-grpc-cart-svc/pkg/pb"
)

type Server struct {
	H          db.Handler
	ProductSvc client.ProductServiceClient
	pb.UnimplementedCartServiceServer
}

type cartResult struct {
	ProductId int64
	Total     int64
	Qty       int64
}

func (s *Server) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.AddToCartResponce, error) {
	// chekc there is a cart is for the user if no create one
	tx := s.H.DB.Begin()
	var cart models.Cart

	findCart := `SELECT id,product_id,total,user_id,qty FROM carts WHERE user_id=$1 AND product_id=$2`
	if err := tx.Raw(findCart, req.UserId, req.ProductId).Scan(&cart).Error; err != nil {
		tx.Rollback()
		return &pb.AddToCartResponce{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}

	fmt.Println("add to cart-cart", cart)

	//find the details of  the product
	product, err := s.ProductSvc.FindOne(req.ProductId)
	if err != nil {
		tx.Rollback()
		return &pb.AddToCartResponce{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}

	fmt.Println("product", product)

	if product.Data.Stock <= 0 {
		tx.Rollback()
		return &pb.AddToCartResponce{
			Status: http.StatusNoContent,
			Error:  "stock too less",
		}, fmt.Errorf("less stock")
	}
	// check if the product is alredy in the cart
	if cart.Id == 0 {
		addProduct := `INSERT INTO carts (user_id, product_id, total, qty) VALUES ($1,$2,$3,$4)`
		if err := tx.Exec(addProduct, req.UserId, req.ProductId, product.Data.Price, 1).Error; err != nil {
			tx.Rollback()
			return &pb.AddToCartResponce{
				Status: http.StatusNoContent,
				Error:  err.Error(),
			}, err
		}
	} else {
		updatCart := `UPDATE carts SET qty=$1,total=$2 WHERE user_id=$3 AND product_id=$4`
		if err := tx.Exec(updatCart, cart.Qty+1, cart.Total+product.Data.Price, req.UserId, req.ProductId).Error; err != nil {
			tx.Rollback()
			return &pb.AddToCartResponce{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			}, err
		}
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return &pb.AddToCartResponce{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}

	return &pb.AddToCartResponce{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) RemoveFromCart(ctx context.Context, req *pb.RemoveFromCartRequest) (*pb.RemoveFromcartResponce, error) {
	//check if the item is present in cart

	fmt.Println("------------", req.ProductId, req.UserId)

	tx := s.H.DB.Begin()
	var cart models.Cart
	findCart := `SELECT * FROM carts WHERE user_id=$1 AND product_id=$2`
	if err := tx.Raw(findCart, req.UserId, req.ProductId).Scan(&cart).Error; err != nil {
		tx.Rollback()
		return &pb.RemoveFromcartResponce{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}
	if cart.Id == 0 {
		tx.Rollback()
		return &pb.RemoveFromcartResponce{
			Status: http.StatusBadRequest,
			Error:  "no product in cart",
		}, fmt.Errorf("no product in cart")
	}
	//get product details
	product, err := s.ProductSvc.FindOne(req.ProductId)
	if err != nil {
		tx.Rollback()
		return &pb.RemoveFromcartResponce{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}
	// if there are multiple items reduce the total qyt
	if cart.Qty == 1 {
		deletItem := `DELETE  FROM carts WHERE product_id=$1 AND user_id=$2`
		if err := tx.Exec(deletItem, req.ProductId, req.UserId).Error; err != nil {
			tx.Rollback()
			return &pb.RemoveFromcartResponce{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			}, err
		}
	} else {
		updateCart := `UPDATE carts set total=$1,qty=$2 WHERE user_id=$3 AND product_id=$4 `
		if err := tx.Exec(updateCart, cart.Total-product.Data.Price, cart.Qty-1, req.UserId, req.ProductId).Error; err != nil {
			tx.Rollback()
			return &pb.RemoveFromcartResponce{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			}, err
		}
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return &pb.RemoveFromcartResponce{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}

	return &pb.RemoveFromcartResponce{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) FindCart(ctx context.Context, req *pb.FindCartRequest) (*pb.FindCartResponse, error) {
	tx := s.H.DB.Begin()
	var carts []*pb.FindCartData
	findCart := `SELECT * FROM carts WHERE user_id=$1`
	if err := tx.Raw(findCart, req.UserId).Scan(&carts).Error; err != nil {
		tx.Rollback()
		return &pb.FindCartResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}
	if len(carts) == 0 {
		tx.Rollback()
		return &pb.FindCartResponse{
			Status: http.StatusBadRequest,
			Error:  "no items in cart",
		}, fmt.Errorf("no items in cart")
	}
	var cartTotal int64
	findCartTotal := `SELECT SUM(total) AS cartTotal FROM carts WHERE user_id = $1;`
	if err := tx.Raw(findCartTotal, req.UserId).Scan(&cartTotal).Error; err != nil {
		tx.Rollback()
		return &pb.FindCartResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return &pb.FindCartResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}
	return &pb.FindCartResponse{
		Status: http.StatusOK,
		Data:   carts,
		Total:  cartTotal,
	}, nil
}

func (s *Server) DeleteCart(ctx context.Context, req *pb.DeleteCartRequest) (*pb.DeleteCartResponse, error) {
	deleteCart := `DELETE FROM carts WHERE user_id=$1`
	if err := s.H.DB.Exec(deleteCart, req.UserId).Error; err != nil {
		return &pb.DeleteCartResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}
	return &pb.DeleteCartResponse{
		Status: http.StatusOK,
	}, nil

}
