package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/guilchaves/treinamento-goexpert/6-APIs/1/configs"
	"github.com/guilchaves/treinamento-goexpert/6-APIs/1/internal/entity"
	"github.com/guilchaves/treinamento-goexpert/6-APIs/1/internal/infra/database"
	"github.com/guilchaves/treinamento-goexpert/6-APIs/1/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/guilchaves/treinamento-goexpert/6-APIs/1/docs"
)

// @title													Go Expert API Example
// @version												1.0
// @description										This is a sample server.
// @termsOfService								http://swagger.io/terms/

// @contact.name									Guilherme Chaves

// @license.name									Full Cycle license
// @license.url										http://www.fullcycle.com.br

// @host													localhost:8000
// @BasePath											/
// @securityDefinitions.apiKey		ApiKeyAuth
// @in														header
// @name													Authorization
func main() {
	configs, err := configs.LoadConfig(".")
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

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", configs.JWTExpiresIn))

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Post("/", productHandler.CreateProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Post("/generate_token", userHandler.GetJWT)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}
