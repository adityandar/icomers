package route

import (
	"icomers/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App               *fiber.App
	UserController    *http.UserController
	ProductController *http.ProductController
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/register", c.UserController.Register)
	c.App.Post("/api/login", c.UserController.Login)
	// change to protected route
	c.App.Post("/api/products", c.ProductController.Create)
	c.App.Get("/api/products", c.ProductController.Get)
}
