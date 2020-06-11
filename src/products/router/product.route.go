package router

import (
	"golang-starter/infrastructure/db"
	"golang-starter/src/products/controllers"
	"golang-starter/src/products/repositories"
	"golang-starter/src/products/services"

	"github.com/gofiber/fiber"
)

func RecipesRoute(app *fiber.App) {
	db := db.Load()
	productRepository := repositories.ProvideProductRepostiory(db)
	productService := services.ProvideProductService(productRepository)
	productController := controllers.ProvideProductController(productService)

	app.Get("/products", productController.GetProducts)
}