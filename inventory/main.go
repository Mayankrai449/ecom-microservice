package main

import (
	"net"

	"github.com/Mayankrai449/ecom-microservice/inventory/db"
	"github.com/Mayankrai449/ecom-microservice/inventory/models"
	pb "github.com/Mayankrai449/ecom-microservice/inventory/proto/inventory"
	"github.com/Mayankrai449/ecom-microservice/inventory/routes"
	service "github.com/Mayankrai449/ecom-microservice/inventory/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
)

func main() {

	db.DB_Config()
	database := db.GetDB()
	db.Migrate(&models.Product{})

	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		pb.RegisterInventoryServiceServer(grpcServer, service.NewInventoryServer(database))

		log.Info("Starting gRPC server on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	routes.RegisterRoutes(e)

	log.Info("Starting HTTP server on :8000")
	e.Logger.Fatal(e.Start(":8000"))

}
