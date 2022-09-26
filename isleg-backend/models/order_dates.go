package models

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

type OrderDates struct {
	ID        uuid.UUID `json:"id"`
	Date      string    `json:"date"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}

type OrderTimes struct {
	ID          uuid.UUID `json:"id"`
	OrderDateID uuid.UUID `json:"order_date_id"`
	Time        string    `json:"time"`
	CreatedAt   string    `json:"-"`
	UpdatedAt   string    `json:"-"`
	DeletedAt   string    `json:"-"`
}

func ValidateOrderDateAndTime(date string, times []string) error {

	hourAndMinute := regexp.MustCompile("([01]?[0-9]|2[0-3]):[0-5][0-9]")

	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return err
	}

	for _, v := range times {

		ts := strings.Split(v, " - ")

		for _, v1 := range ts {

			checkHour := hourAndMinute.MatchString(v1)
			if !checkHour {
				return errors.New("the data type should be hour")
			}

		}

	}

	return nil

}
