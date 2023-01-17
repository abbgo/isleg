package pkg

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// mulitpart file upload
func MultipartFileUpload(nameUploadedFile, path string, context *gin.Context) ([]string, error) {

	var fileNames []string

	files := context.Request.MultipartForm.File[nameUploadedFile]

	for _, v := range files {

		extension := filepath.Ext(v.Filename)
		//validate image
		if extension != ".jpg" && extension != ".JPG" && extension != ".jpeg" && extension != ".JPEG" && extension != ".png" && extension != ".PNG" && extension != ".gif" && extension != ".GIF" && extension != ".svg" && extension != ".SVG" && extension != ".WEBP" && extension != ".webp" {
			return []string{}, errors.New("the file must be an image")
		}
		fileName := uuid.New().String() + extension

		if err := context.SaveUploadedFile(v, ServerPath+"/uploads/"+path+"/"+fileName); err != nil {
			return []string{}, err
		}

		fileNames = append(fileNames, "uploads/"+path+"/"+fileName)

	}
	return fileNames, nil

}

// File upload
func FileUpload(fileName, path string, context *gin.Context) (string, error) {

	file, err := context.FormFile(fileName)
	if err != nil {
		return "", err
	}

	extensionFile := filepath.Ext(file.Filename)

	// VALIDATE IMAGE
	if extensionFile != ".jpg" && extensionFile != ".JPG" && extensionFile != ".jpeg" && extensionFile != ".JPEG" && extensionFile != ".png" && extensionFile != ".PNG" && extensionFile != ".gif" && extensionFile != ".GIF" && extensionFile != ".svg" && extensionFile != ".SVG" && extensionFile != ".WEBP" && extensionFile != ".webp" {
		return "", errors.New("the file must be an image")
	}

	newFileName := uuid.New().String() + extensionFile
	if err := context.SaveUploadedFile(file, ServerPath+"uploads/"+path+"/"+newFileName); err != nil {
		return "", err
	}

	return "uploads/" + path + "/" + newFileName, nil

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
		if extensionFile != ".jpg" && extensionFile != ".JPG" && extensionFile != ".jpeg" && extensionFile != ".JPEG" && extensionFile != ".png" && extensionFile != ".PNG" && extensionFile != ".gif" && extensionFile != ".GIF" && extensionFile != ".svg" && extensionFile != ".SVG" && extensionFile != ".WEBP" && extensionFile != ".webp" {
			return "", errors.New("the file must be an image")
		}

		newFileName = uuid.New().String() + extensionFile
		if err := context.SaveUploadedFile(file, ServerPath+"uploads/"+path+"/"+newFileName); err != nil {
			return "", err
		}

		if err := os.Remove(ServerPath + oldFileName); err != nil {
			return "", err
		}

		return "uploads/" + path + "/" + newFileName, nil

	}

}
