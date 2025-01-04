package entity

type Product struct {
	ID          int     `gorm:"uniqueKey"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	CreatedAt   int64   `gorm:"type:bigint;autoCreateTime:milli"`
	UpdatedAt   int64   `gorm:"type:bigint;autoUpdateTime:milli"`
}
