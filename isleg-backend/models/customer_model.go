package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type Customer struct {
	ID          uuid.UUID      `json:"id"`
	FullName    string         `json:"full_name"`
	PhoneNumber string         `json:"phone_number"`
	Password    string         `json:"password"`
	Birthday    time.Time      `json:"birthday"`
	Gender      string         `json:"gender"`
	Addresses   pq.StringArray `json:"addresses"`
	CreatedAt   string         `json:"-"`
	UpdatedAt   string         `json:"-"`
	DeletedAt   string         `json:"-"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	hashPassword := string(bytes)
	return hashPassword, nil
}

func CheckPassword(providedPassword, oldPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(oldPassword), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
