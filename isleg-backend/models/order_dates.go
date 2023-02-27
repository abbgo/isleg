package models

import (
	"errors"

	"github.com/google/uuid"
)

type DateHours struct {
	ID        string        `json:"id,omitempty"`
	Hour      uint          `json:"hour,omitempty" binding:"required"`
	DateID    uuid.NullUUID `json:"date_id,omitempty" binding:"required"`
	CreatedAt string        `json:"-"`
	UpdatedAt string        `json:"-"`
	DeletedAt string        `json:"-"`
}

type OrderDates struct {
	ID                    string                  `json:"id,omitempty"`
	Date                  string                  `json:"date,omitempty" binding:"required"`
	CreatedAt             string                  `json:"-"`
	UpdatedAt             string                  `json:"-"`
	DeletedAt             string                  `json:"-"`
	TranslationOrderDates []TranslationOrderDates `json:"translation_order_dates,omitempty" binding:"required"` // one to many
}

type TranslationOrderDates struct {
	ID          string        `json:"id,omitempty"`
	LangID      uuid.NullUUID `json:"lang_id,omitempty" binding:"required"`
	OrderDateID uuid.NullUUID `json:"order_date_id,omitempty"`
	Date        string        `json:"date,omitempty" binding:"required"`
	CreatedAt   string        `json:"-"`
	UpdatedAt   string        `json:"-"`
	DeletedAt   string        `json:"-"`
}

type DateHourTimes struct {
	ID         string        `json:"id,omitempty"`
	DateHourID uuid.NullUUID `json:"date_hour_id,omitempty" binding:"required"`
	TimeID     uuid.NullUUID `json:"time_id,omitempty" binding:"required"`
	CreatedAt  string        `json:"-"`
	UpdatedAt  string        `json:"-"`
	DeletedAt  string        `json:"-"`
}

type OrderTimes struct {
	ID        string `json:"id,omitempty"`
	Time      string `json:"time,omitempty" binding:"required"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}

func ValidateOrderDateAndTime(date string) error {

	if date != "today" && date != "tomorrow" {
		return errors.New("the date should be today or tomorrow")
	}

	// hourAndMinute := regexp.MustCompile("([01]?[0-9]|2[0-3]):[0-5][0-9]")
	// for _, v := range times {

	// 	ts := strings.Split(v.Time, " - ")

	// 	for _, v1 := range ts {

	// 		checkHour := hourAndMinute.MatchString(v1)
	// 		if !checkHour {
	// 			return errors.New("the data type should be hour")
	// 		}

	// 	}

	// }

	return nil

}
