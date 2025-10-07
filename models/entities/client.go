package entities

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name       string
	FamilyName string
}
