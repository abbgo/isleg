package helpers

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// File upload
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
	if err := context.SaveUploadedFile(file, "./uploads/"+path+"/"+newFileName); err != nil {
		return "", err
	}

	return newFileName, nil

}

// File upload for update function
func FileUploadForUpdate(fileName, path, oldFileName string, context *gin.Context) (string, error) {

	var newFileName string

	file, err := context.FormFile(fileName)
	if err != nil {

		newFileName = oldFileName
		return newFileName, nil

	} else {

		extensionFile := filepath.Ext(file.Filename)

		// VALIDATE IMAGE
		if extensionFile != ".jpg" && extensionFile != ".jpeg" && extensionFile != ".png" && extensionFile != ".gif" {
			return "", errors.New("the file must be an image")
		}

		newFileName = uuid.New().String() + extensionFile
		if err := context.SaveUploadedFile(file, "./uploads/"+path+"/"+newFileName); err != nil {
			return "", err
		}

		if err := os.Remove("./" + oldFileName); err != nil {
			return "", err
		}

		return newFileName, nil

	}

}
