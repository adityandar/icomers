package model

type UserResponse struct {

	// ID        int    `gorm:"uniqueKey"`
	// Username  string `gorm:"unique; not null"`
	// Email     string `gorm:"unique: not null"`
	// Password  string `gorm:"not null"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`

	ID        int    `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"email,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}
