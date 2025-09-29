package models

type City struct {
	BaseModel
	Name      string  `gorm:"type:string;not null"`
	CountryID int     `gorm:"not null"`
	Country   Country `gorm:"foreignKey:CountryID"`
}
