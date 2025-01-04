package config

import (
	"icomers/internal/delivery/http"
	"icomers/internal/delivery/route"
	"icomers/internal/repository"
	"icomers/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootStrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootStrapConfig) {
	// setup repositories
	userRepository := repository.NewUserRepository(config.Log, config.Config)
	productRepository := repository.NewProductRepository(config.Log)
	// setup usecases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)
	productUseCase := usecase.NewProductUsecase(config.DB, config.Log, config.Validate, productRepository)
	// setup controller
	userController := http.NewUserController(config.Log, userUseCase)
	productController := http.NewProductController(config.Log, productUseCase)
	// setup middleware

	// setup route
	routeConfig := route.RouteConfig{
		App:               config.App,
		UserController:    userController,
		ProductController: productController,
	}
	routeConfig.Setup()
}
