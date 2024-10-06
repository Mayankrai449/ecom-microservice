package main

import (
	"github.com/Mayankrai449/ecom-microservice/orders/db"
	"github.com/Mayankrai449/ecom-microservice/orders/models"
	pb "github.com/Mayankrai449/ecom-microservice/orders/proto/inventory"
	"github.com/Mayankrai449/ecom-microservice/orders/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
)

func main() {

	db.DB_Config()
	database := db.GetDB()
	db.Migrate(&models.Order{})

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic("failed to connect to inventory service: " + err.Error())
	}
	defer conn.Close()
	inventoryClient := pb.NewInventoryServiceClient(conn)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	routes.RegisterRoutes(e, database, inventoryClient)

	e.Logger.Fatal(e.Start(":8080"))

}
