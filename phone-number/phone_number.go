package phonenumber

import (
	"fmt"
	"regexp"
	"strings"
)

func getNumberParts(ph string) (areaCode string, exchCode string, subNbr string) {
	return ph[:3], ph[3:6], ph[6:]

}

func Number(phoneNumber string) (string, error) {
	r := regexp.MustCompile(`\D`)
	s := r.ReplaceAllString(phoneNumber, "")
	s = strings.TrimLeft(s, "1")
	if len(s) != 10 {
		return "", fmt.Errorf("not 10 digits")
	}
	areaCode, exchCode, _ := getNumberParts(s)

	if areaCode[0] == '0' {
		return "", fmt.Errorf("area code cannot start with 0")
	}
	if exchCode[0] == '0' || exchCode[0] == '1' {
		return "", fmt.Errorf("exchange code cannot start with 0 or 1")
	}

	return s, nil
}

func AreaCode(phoneNumber string) (string, error) {
	ph, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	areaCode, _, _ := getNumberParts(ph)
	return areaCode, nil
}

func Format(phoneNumber string) (string, error) {
	ph, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	areaCode, exchCode, subNbr := getNumberParts(ph)
	fph := fmt.Sprintf("(%s) %s-%s", areaCode, exchCode, subNbr)
	return fph, nil
}
