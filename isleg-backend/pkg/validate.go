package pkg

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func ValidateMiddlewareData(c *gin.Context, value string) (string, error) {
	ID, hasCustomer := c.Get(value)
	if !hasCustomer {
		return "", errors.New(value + " is required")
	}
	id, ok := ID.(string)
	if !ok {
		return "", errors.New(value + "id must be string")
	}

	return id, nil
}
