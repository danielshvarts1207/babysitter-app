package entities

import "gorm.io/gorm"

type Role string

const (
	BABYSITTER Role = "BABYSITTER"
	CLIENT     Role = "CLIENT"
	ADMIN      Role = "ADMIN"
)

type User struct {
	gorm.Model
	Name         string
	FamilyName   string
	Email        string
	Role         Role `gorm:"type:enum('BABYSITTER', 'CLIENT', 'ADMIN');column:role"`
	PasswordHash string
}
