package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {

	// regex to convert string to alphanumeric
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

	cleanString := reg.ReplaceAllString(s, "")
	cleanString = strings.ToLower(cleanString)

	length := len([]rune(cleanString))

	for index, i := range cleanString {
		if string(i) != string(cleanString[length-index-1]) {
			return false
		}
	}

	return true
}
