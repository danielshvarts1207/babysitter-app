package dtos

type CreateUserDto struct {
	Name         string `json:"name"  binding:"required"`
	FamilyName   string `json:"familyName"  binding:"required"`
	Email        string `json:"email"  binding:"required"`
	Role         string `json:"role"  binding:"required,oneof=BABYSITTER CLIENT ADMIN"`
	PasswordHash string `json:"passwordHash"  binding:"required"`
}
