package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Converter(s string) (string, error) {
	Text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var res string

	if strings.ContainsAny(s, Text) {
		res = morse.ToText(s)
	} else {
		res = morse.ToMorse(s)
	}

	return res, nil
}
