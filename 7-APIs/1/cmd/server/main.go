package main

import (
	"net/http"

	"github.com/guilchaves/treinamento-goexpert/6-APIs/1/configs"
	"github.com/guilchaves/treinamento-goexpert/6-APIs/1/internal/entity"
	"github.com/guilchaves/treinamento-goexpert/6-APIs/1/internal/infra/database"
	"github.com/guilchaves/treinamento-goexpert/6-APIs/1/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDb := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDb)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)
}
