package utils

import (
	"fmt"
	"net/http"
	"net/mail"
	"regexp"
	"strings"
)

var (
	MsgInvalidEmail = "Write correct email"
	MsgInvalidName  = "Write correct name. Username should start with an alphabet [A-Za-z] and all other characters can be alphabets, numbers or an underscore so, [A-Za-z0-9_]. The username consists of 5 to 15 characters inclusive."
	MsgInvalidPass  = "Password must contain letters, numbers and must be at least 6 characters."
)

const maxImageSizeBytes = 1048576 // 1MB

func IsValidEmail(email string) bool {
	rxEmail := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if len(email) > 254 || !rxEmail.MatchString(email) {
		return false
	}
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsValidPassword(pass string) bool {
	return len(pass) >= 6
}

func IsValidImageURL(imageURL string) bool {
	response, err := http.Get(imageURL)
	if err != nil {
		return false
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println(response.StatusCode)
		return false
	}

	contentType := response.Header.Get("Content-Type")
	validTypes := []string{"image/jpeg", "image/png", "image/gif", "image/bmp"}
	for i, validType := range validTypes {
		if contentType == validType {
			break
		}
		if i == len(validTypes)-1 {
			return false
		}
	}

	return response.ContentLength <= maxImageSizeBytes

}

func IsValidPrice(price float64) bool {
	return price >= 0
}

func IsValidTitle(title string) bool {
	title = strings.TrimSpace(title)
	if len(title) < 5 || len(title) > 100 {
		return false
	}
	return true
}

func IsValidDescription(description string) bool {
	description = strings.TrimSpace(description)
	if len(description) < 5 || len(description) > 1000 {
		return false
	}
	return true
}
