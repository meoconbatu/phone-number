package phonenumber

import (
	"errors"
	"strings"
)

const testVersion = 2

var formatPhoneNumber = "(XXX) XXX-XXXX"
var phoneNumberLength = 10

func Number(input string) (string, error) {
	phoneNumber := make([]rune, phoneNumberLength)
	index := phoneNumberLength - 1
	for i := len(input) - 1; i >= 0; i-- {
		n := rune(input[i])
		isNumber := n >= '0' && n <= '9'
		isLetter := n >= 'a' && n <= 'z'
		if isNumber {
			if index < -1 {
				return "", errors.New("invalid when more than 11 digits.")
			} else if index == -1 {
				if n != '1' {
					return "", errors.New("invalid when 11 digits does not start with a 1.")
				}
			} else {
				if n == '0' || n == '1' {
					if index == 3 {
						return "", errors.New("invalid if exchange code does not start with 2-9.")
					} else if index == 0 {
						return "", errors.New("invalid if area code does not start with 2-9.")
					}
				}
				phoneNumber[index] = n
				index--
			}
		} else if isLetter {
			return "", errors.New("invalid with right number of digits but letters mixed in.")
		}
	}
	if index >= 0 {
		return "", errors.New("invalid when 9 digits.")
	}
	return string(phoneNumber), nil
}
func AreaCode(input string) (string, error) {
	n, err := Number(input)
	if err == nil {
		return n[0:3], nil
	}
	return "", err
}
func Format(input string) (string, error) {
	n, err := Number(input)
	output := formatPhoneNumber
	if err != nil {
		return "", err
	}
	for _, nn := range n {
		output = strings.Replace(output, "X", string(nn), 1)
	}
	return output, nil
}
