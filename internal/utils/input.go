package utils

import (
	"fmt"
	"strings"
)

func GetUserInput(validAnswers ...string) string {
	var choice string
	var isValid bool = false

	for !isValid {
		fmt.Scanln(&choice)
		for _, v := range validAnswers {
			if strings.EqualFold(v, choice) {
				isValid = true
			}
		}
	}
	return strings.ToUpper(choice)
}
