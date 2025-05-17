package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func mors(str string) bool {
	b := strings.Contains(str, ".-")
	return b
}

func Convert(in string) (string, error) {

	if mors(in) == true {
		result := morse.ToText(in)

		if len(strings.Split(result, "")) == 0 {
			return "", errors.New("ошибка распознавания кода Морзе")
		}
		return result, nil
	}
	result := morse.ToMorse(in)
	if len(strings.Split(result, "")) == 0 {
		return "", errors.New("ошибка распознавания текста")
	}
	return result, nil
}
