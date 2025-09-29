package models

type Country struct {
	BaseModel
	Name string `gorm:"type:string;not null"`
	City []City
}
