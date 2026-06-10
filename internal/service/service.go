package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Converter(s string) (string, error) {
	Text := "АаБбВвГгДдЕеЁёЖжЗзИиЙйКкЛлМмНнОоПпРрСсТтУуФфХхЦцЧчШшЩщЪъЫыЬьЭэЮюЯя"
	var res string

	if strings.ContainsAny(s, Text) {
		res = morse.ToMorse(s)
	} else {
		res = morse.ToText(s)
	}

	return res, nil
	
}
