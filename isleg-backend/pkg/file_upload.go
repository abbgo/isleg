package pkg

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// mulitpart file upload
func MultipartFileUpload(nameUploadedFile, path string, context *gin.Context) ([]string, error) {

	var fileNames []string

	form, _ := context.MultipartForm()

	// Retrieve the files
	files := form.File[nameUploadedFile]

	_, err := os.Stat(ServerPath + "uploads/" + path)
	if err != nil {
		if err := os.MkdirAll(ServerPath+"uploads/"+path, os.ModePerm); err != nil {
			return []string{}, err
		}
	}

	for _, v := range files {

		imageName := v.Filename
		extension := filepath.Ext(imageName)
		//validate image
		if extension != ".jpg" && extension != ".JPG" && extension != ".jpeg" && extension != ".JPEG" && extension != ".png" && extension != ".PNG" && extension != ".gif" && extension != ".GIF" && extension != ".svg" && extension != ".SVG" && extension != ".WEBP" && extension != ".webp" {
			return []string{}, errors.New("the file must be an image")
		}
		// fileName := uuid.New().String() + extension
		if strings.Contains(imageName, " ") {
			imageName = strings.ReplaceAll(imageName, " ", "_")
		}
		fileName := imageName

		if err := context.SaveUploadedFile(v, ServerPath+"uploads/"+path+"/"+fileName); err != nil {
			return []string{}, err
		}

		fileNames = append(fileNames, "uploads/"+path+"/"+fileName)

	}
	return fileNames, nil

}

// File upload
func FileUpload(fileName, path, fileType string, context *gin.Context) (string, error) {

	file, err := context.FormFile(fileName)
	if err != nil {
		return "", err
	}

	extensionFile := filepath.Ext(file.Filename)

	var newFileName string

	// VALIDATE IMAGE
	if fileType == "image" {
		if extensionFile != ".jpg" && extensionFile != ".JPG" && extensionFile != ".jpeg" && extensionFile != ".JPEG" && extensionFile != ".png" && extensionFile != ".PNG" && extensionFile != ".gif" && extensionFile != ".GIF" && extensionFile != ".svg" && extensionFile != ".SVG" && extensionFile != ".WEBP" && extensionFile != ".webp" {
			return "", errors.New("the file must be an image")
		}
		newFileName = uuid.New().String() + extensionFile

	} else if fileType == "excel" {
		if extensionFile != ".xlsx" && extensionFile != ".xlsm" && extensionFile != ".xlsb" && extensionFile != ".xltx" {
			return "", errors.New("the file must be an excel")
		}
		newFileName = "products" + extensionFile
	} else {
		return "", errors.New("invalid file type")
	}

	_, err = os.Stat(ServerPath + "uploads/" + path)
	if err != nil {
		if err := os.MkdirAll(ServerPath+"uploads/"+path, os.ModePerm); err != nil {
			return "", err
		}
	}
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
