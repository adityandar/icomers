package entity

type User struct {
	ID        int    `gorm:"uniqueKey"`
	Username  string `gorm:"unique; not null"`
	Email     string `gorm:"unique: not null"`
	Password  string `gorm:"not null"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
}
