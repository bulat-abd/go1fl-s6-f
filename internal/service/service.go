package service

import (
    "strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(input string) bool{
	r := strings.NewReplacer(" ", "", ".", "","-","")
	return (len(r.Replace(input))==0)
}

func Transcode (input string) string {
    if isMorse(input) {
		return morse.ToText(input)
	} else {
		return morse.ToMorse(input)
	}
}