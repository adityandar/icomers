package dto

import (
	"icomers/models"
	"time"
)

func ConvertToUserResponse(user models.User) UserResponse {
	var deletedAt *time.Time

	if user.DeletedAt.Valid {
		deletedAt = &user.DeletedAt.Time
	}

	return UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: deletedAt,
	}
}
