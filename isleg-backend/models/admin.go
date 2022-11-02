package models

type Admin struct {
	ID          string `json:"id,omitempty"`
	FullName    string `json:"full_name,omitempty" binding:"required,min=3"`
	PhoneNumber string `json:"phone_number,omitempty" binding:"required,e164,len=12"`
	Password    string `json:"password,omitempty" binding:"required,min=5,max=25"`
	CreatedAt   string `json:"-"`
	UpdatedAt   string `json:"-"`
	DeletedAt   string `json:"-"`
}
