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
	userRepository := repository.NewUserRepository(config.Log)
	// setup usecases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)
	// setup controller
	userController := http.NewUserController(userUseCase, config.Log)
	// setup middleware

	// setup route
	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
	}
	routeConfig.Setup()
}
