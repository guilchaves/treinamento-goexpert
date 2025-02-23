package main

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/chi/v5"
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

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/products", productHandler.GetProducts)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Post("/products", productHandler.CreateProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	http.ListenAndServe(":8000", r)
}
