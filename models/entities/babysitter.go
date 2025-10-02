package entities

import "gorm.io/gorm"

type Babysitter struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	Name       string
	FamilyName string
	HourPrice  float64
}
