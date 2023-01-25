package models

import (
	"errors"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

type OrderDates struct {
	ID                    string                  `json:"id,omitempty"`
	Date                  string                  `json:"date,omitempty" binding:"required"`
	CreatedAt             string                  `json:"-"`
	UpdatedAt             string                  `json:"-"`
	DeletedAt             string                  `json:"-"`
	TranslationOrderDates []TranslationOrderDates `json:"translation_order_dates,omitempty" binding:"required"` // one to many
	OrderTimes            []OrderTimes            `json:"order_times,omitempty" binding:"required"`             // one to many
}

type OrderTimes struct {
	ID          string        `json:"id,omitempty"`
	OrderDateID uuid.NullUUID `json:"order_date_id,omitempty"`
	Time        string        `json:"time,omitempty" binding:"required"`
	CreatedAt   string        `json:"-"`
	UpdatedAt   string        `json:"-"`
	DeletedAt   string        `json:"-"`
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

func ValidateOrderDateAndTime(date string, times []OrderTimes) error {

	hourAndMinute := regexp.MustCompile("([01]?[0-9]|2[0-3]):[0-5][0-9]")

	if date != "today" && date != "tomorrow" {
		return errors.New("the date should be today or tomorrow")
	}

	for _, v := range times {

		ts := strings.Split(v.Time, " - ")

		for _, v1 := range ts {

			checkHour := hourAndMinute.MatchString(v1)
			if !checkHour {
				return errors.New("the data type should be hour")
			}

		}

	}

	return nil

}
