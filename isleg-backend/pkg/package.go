package pkg

import (
	"errors"
	"path/filepath"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// email validate
func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func FileUpload(fileName, path string, context *gin.Context) (string, error) {

	file, err := context.FormFile(fileName)
	if err != nil {
		return "", err
	}

	extensionFile := filepath.Ext(file.Filename)

	// VALIDATE IMAGE
	if extensionFile != ".jpg" && extensionFile != ".jpeg" && extensionFile != ".png" && extensionFile != ".gif" {
		return "", errors.New("the file must be an image")
	}

	newFileName := uuid.New().String() + extensionFile
	context.SaveUploadedFile(file, "./uploads/"+path+"/"+newFileName)

	return newFileName, nil

}
