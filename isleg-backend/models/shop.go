package models

type Shop struct {
	ID          string `json:"id,omitempty"`
	OwnerName   string `json:"owner_name,omitempty" binding:"required"`
	Address     string `json:"address,omitempty" binding:"required"`
	PhoneNumber string `json:"phone_number,omitempty" binding:"required"`
	RunningTime string `json:"running_time,omitempty" binding:"required"`
	CreatedAt   string `json:"-"`
	UpdatedAt   string `json:"-"`
	DeletedAt   string `json:"-"`
}
