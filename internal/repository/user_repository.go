package repository

import (
	"fmt"
	"icomers/internal/entity"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
	Log    *logrus.Logger
	Config *viper.Viper
}

func NewUserRepository(log *logrus.Logger, config *viper.Viper) *UserRepository {
	return &UserRepository{
		Log:    log,
		Config: config,
	}
}

func (r *UserRepository) FindByToken(db *gorm.DB, user *entity.User, token string) error {
	return db.Where("token = ?", token).First(user).Error
}

func (r *UserRepository) CountByUsername(db *gorm.DB, username string) (int64, error) {
	var total int64
	err := db.Model(new(entity.User)).Where("username = ?", username).Count(&total).Error

	return total, err
}

func (r *UserRepository) FindByEmail(db *gorm.DB, user *entity.User, email string) error {
	return db.Where("email = ?", email).First(user).Error
}

func (r *UserRepository) CreateToken(user *entity.User) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		Issuer:    fmt.Sprintf("%d", user.ID),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	r.Log.Warnf("secret key : %+v", r.Config.GetString("jwt.secret_key"))
	secretKey := r.Config.GetString("jwt.secret_key")
	return token.SignedString([]byte(secretKey))
}
