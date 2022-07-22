package main

import (
	"cashier/handler"
	"cashier/product"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/cashier?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Success connecting to the database")

	productRepository := product.NewRepository(db)

	productService := product.NewService(productRepository)

	productHandler := handler.NewProductHandler(productService)

	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")

	api.GET("/products", productHandler.GetProducts)
	api.POST("/products", productHandler.CreateProduct)

	api.GET("/products/:id", productHandler.GetProduct)
	api.PUT("/products/:id", productHandler.UpdateProduct)

	api.GET("/prices", productHandler.GetPrices)
	api.POST("/prices", productHandler.CreatePrice)

	api.GET("/prices/:id", productHandler.GetPrice)
	api.PUT("/prices/:id", productHandler.UpdatePrice)

	router.Run() // listen and serve on 0.0.0.0:8080
}
