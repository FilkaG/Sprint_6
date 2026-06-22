package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Detect(data string) (string, error) {
	if data == "" {
		return "", errors.New("empty data")
	}

	trimmed := strings.TrimFunc(data, func(r rune) bool {
		return r == '.' || r == '-' || unicode.IsSpace(r)
	})

	if trimmed == "" {
		return morse.ToText(data), nil
	}

	return morse.ToMorse(data), nil
}
