package service

import (
	"context"

	pb "github.com/Mayankrai449/ecom-microservice/inventory/proto/inventory"

	"github.com/Mayankrai449/ecom-microservice/inventory/models"
	"gorm.io/gorm"
)

type InventoryServer struct {
	pb.UnimplementedInventoryServiceServer
	db *gorm.DB
}

func NewInventoryServer(db *gorm.DB) *InventoryServer {
	return &InventoryServer{db: db}
}

func (s *InventoryServer) CheckStock(ctx context.Context, req *pb.StockRequest) (*pb.StockResponse, error) {
	for i, productID := range req.ProductIds {
		var product models.Product
		if err := s.db.First(&product, productID).Error; err != nil {
			return &pb.StockResponse{
				InStock: false,
				Message: "Product not found",
			}, nil
		}

		if uint32(product.Stock) < req.Quantities[i] {
			return &pb.StockResponse{
				InStock: false,
				Message: "Insufficient stock",
			}, nil
		}
	}

	return &pb.StockResponse{
		InStock: true,
		Message: "All products in stock",
	}, nil
}
