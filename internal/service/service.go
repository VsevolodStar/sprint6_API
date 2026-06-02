package service

import (
	"errors"
	"unicode"

	"sprint6_API/pkg/morse"
)

const content = "заглушка" //переделать на входящий файл

func isMorse(content string) bool {
	for _, char := range content {
		if char != '.' && char != '-' && char != ' ' {
			return false
		}
	}
	return true
}

func isText(content string) bool {
	for _, char := range content {
		if unicode.Is(unicode.Cyrillic, char) {
			return true
		}
	}
	return false
}

func CheckContentType(content string) string {

	if isMorse(content) == true {
		return "Morse"
	} else if isText(content) == true {
		return "Text"
	} else {
		return "Incorrect format"
	}
}

func ConvertText(content string) (string, error) {

	if CheckContentType(content) == "Morse" {
		return morse.ToText(content), nil
	} else if CheckContentType(content) == "Text" {
		return morse.ToMorse(content), nil
	} else {
		return "", errors.New("Incorrect format. Try to load another file")
	}

}
