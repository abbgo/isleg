package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func SendOTPSmsCode(phoneNumber string) error {
	secretOtp := os.Getenv("JWT_SECRET_KEY")

	// Generate a new TOTP key.
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "isleg Online Market",
		AccountName: phoneNumber,
		Secret:      []byte(secretOtp),
		Algorithm:   otp.AlgorithmSHA1, // You can choose a different algorithm if needed.
		SecretSize:  4,
		Period:      60,
	})
	if err != nil {
		return err
	}

	// Get the current OTP code.
	otpCode, err := totp.GenerateCode(key.Secret(), time.Now())
	if err != nil {
		return err
	}

	// Define the URL you want to send the request to
	url := "localhost:5050"

	// Define the request body as a Go struct
	requestData := struct {
		Phone   string `json:"phone"`
		Content string `json:"content"`
	}{
		Phone:   phoneNumber,
		Content: "isleg.com.tm kot: " + otpCode,
	}

	// Convert the request body to a JSON string
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	// Create a new HTTP POST request with the request body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	// Set the request headers (optional)
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return errors.New("request failed with status code")
	}

	return nil
}
