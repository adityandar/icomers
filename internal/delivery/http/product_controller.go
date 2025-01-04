package http

import (
	"icomers/internal/model"
	"icomers/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ProductController struct {
	Log            *logrus.Logger
	ProductUseCase *usecase.ProductUsecase
}

func NewProductController(log *logrus.Logger, productUseCase *usecase.ProductUsecase) *ProductController {
	return &ProductController{
		Log:            log,
		ProductUseCase: productUseCase,
	}
}

func (c *ProductController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateProductRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse body request : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.ProductUseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create new product %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ProductResponse]{Data: response})
}

func (c *ProductController) Get(ctx *fiber.Ctx) error {
	response, err := c.ProductUseCase.Get(ctx.UserContext())

	if err != nil {
		c.Log.Warnf("Failed to fetch products %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.ProductResponse]{
		Data: response,
	})
}
