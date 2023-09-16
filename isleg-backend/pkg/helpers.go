package pkg

import (
	"io"
	"math"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

var ServerPath = os.Getenv("SERVER_PATH")

func CopyFile(sourceFilePath, destinationFilePath, fileName string) (string, error) {

	// newFileName := uuid.New().String() + strings.Split(fileName, ".")[1]
	newFileName := uuid.New().String() + filepath.Ext(fileName)

	// Open the source file for reading
	source, err := os.Open(ServerPath + sourceFilePath + fileName)
	if err != nil {
		return "", err
	}

	// Create the destination file
	destination, err := os.Create(ServerPath + destinationFilePath + newFileName)
	if err != nil {
		return "", err
	}

	// Copy the contents from source to destination
	_, err = io.Copy(destination, source)
	if err != nil {
		return "", err
	}

	source.Close()
	destination.Close()

	err = os.Remove(ServerPath + sourceFilePath + fileName)
	if err != nil {
		return "", err
	}

	return newFileName, nil

}
