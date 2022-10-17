package models

import (
	"errors"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderDates struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Date      string    `json:"date,omitempty"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}

type OrderTimes struct {
	ID          uuid.UUID `json:"id,omitempty"`
	OrderDateID uuid.UUID `json:"order_date_id,omitempty"`
	Time        string    `json:"time,omitempty"`
	CreatedAt   string    `json:"-"`
	UpdatedAt   string    `json:"-"`
	DeletedAt   string    `json:"-"`
}

type TranslationOrderDates struct {
	ID          uuid.UUID `json:"id,omitempty"`
	LangID      uuid.UUID `json:"lang_id,omitempty"`
	OrderDateID uuid.UUID `json:"order_date_id,omitempty"`
	Date        string    `json:"date,omitempty"`
	CreatedAt   string    `json:"-"`
	UpdatedAt   string    `json:"-"`
	DeletedAt   string    `json:"-"`
}

func ValidateOrderDateAndTime(date string, times []string, languages []Language, dataNames []string, context *gin.Context) error {

	hourAndMinute := regexp.MustCompile("([01]?[0-9]|2[0-3]):[0-5][0-9]")

	// _, err := time.Parse("2006-01-02", date)
	// if err != nil {
	// 	return err
	// }

	if date != "today" && date != "tomorrow" {
		return errors.New("the date should be today or tomorrow")
	}

	for _, dataName := range dataNames {
		for _, v := range languages {
			if context.PostForm(dataName+"_"+v.NameShort) == "" {
				return errors.New(dataName + "_" + v.NameShort + " is required")
			}
		}
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
