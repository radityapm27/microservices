package utils

import (
	"strconv"
)

// StringToInt ...
func (u *utils) StringToInt(input string) (int, error) {
	input = replaceEmptyStringToDigit(input)
	resultFloat, err := strconv.ParseFloat(input, 64)

	if err != nil {
		return 0, err
	}

	result := int(resultFloat)

	return result, nil
}

func replaceEmptyStringToDigit(input string) string {
	if input == "" {
		input = "0"
	}

	return input
}
