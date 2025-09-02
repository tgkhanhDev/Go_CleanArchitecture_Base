package entities

type Role struct {
	ID          string `gorm:"size:20;primaryKey"`
	Description string
	IsMandatory bool
}
