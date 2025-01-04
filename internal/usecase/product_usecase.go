package usecase

import (
	"context"
	"icomers/internal/entity"
	"icomers/internal/model"
	"icomers/internal/model/converter"
	"icomers/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductUsecase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	ProductRepository *repository.ProductRepository
}

func NewProductUsecase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, productRepository *repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{
		DB:                db,
		Log:               log,
		Validate:          validate,
		ProductRepository: productRepository,
	}
}

// create butuh payload createnya
func (c *ProductUsecase) Create(ctx context.Context, request *model.CreateProductRequest) (*model.ProductResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	product := &entity.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
	}

	if err := c.ProductRepository.Create(tx, product); err != nil {

		c.Log.Warnf("Failed to insert user to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ProductToResponse(product), nil
}

func (c *ProductUsecase) Get(ctx context.Context) ([]model.ProductResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	products, err := c.ProductRepository.FindAll(tx)

	if err != nil {
		c.Log.Warnf("Failed to fetch products : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]model.ProductResponse, len(products))

	for i, product := range products {
		responses[i] = *converter.ProductToResponse(&product)
	}

	return responses, nil
}
