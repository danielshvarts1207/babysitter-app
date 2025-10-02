package dtos

type CreateBabysitterDto struct {
	Name       string  `json:"name"  binding:"required"`
	FamilyName string  `json:"familyName"  binding:"required"`
	HourPrice  float64 `json:"hourPrice"  binding:"required"`
}
