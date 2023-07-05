package pkg

import (
	"io"
	"math"
	"os"
	"strings"

	"github.com/google/uuid"
)

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

var ServerPath = os.Getenv("SERVER_PATH")

func CopyFile(sourceFilePath, destinationFilePath, fileName string) (string, error) {

	newFileName := uuid.New().String() + strings.Split(fileName, ".")[1]

	// Open the source file for reading
	source, err := os.Open(ServerPath + sourceFilePath + fileName)
	if err != nil {
		return "", err
	}
	defer source.Close()

	// Create the destination file
	destination, err := os.Create(ServerPath + destinationFilePath + newFileName)
	if err != nil {
		return "", err
	}
	defer destination.Close()

	// Copy the contents from source to destination
	_, err = io.Copy(destination, source)
	if err != nil {
		return "", err
	}

	err = os.Remove(ServerPath + sourceFilePath + fileName)
	if err != nil {
		return "", err
	}

	return newFileName, nil

}
