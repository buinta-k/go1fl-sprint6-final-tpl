package service

import (
	"strings"
)

func main() {
	Text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if strings.ContainsAny(s, Text) {
		res := ToText(s)
	} else {
		res := ToMorse(s)
	}
}
