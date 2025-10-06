package entities

import "gorm.io/gorm"

type Babysitter struct {
	gorm.Model
	Name       string
	FamilyName string
	HourPrice  float64
}
