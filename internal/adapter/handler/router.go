package handler

import (
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	*fiber.App
}

func NewRouter(app *fiber.App, handler ProductHandler) {
	app.Post("/product", handler.CreateProduct)
	app.Get("/product/:id", handler.GetProduct)
	app.Get("/products", handler.GetProducts)
	app.Put("/product/:id", handler.UpdateProduct)
	app.Delete("/product/:id", handler.DeleteProduct)
}
