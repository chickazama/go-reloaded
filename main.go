package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	alphaOffset = 'a' - 'A'
	expArgc     = 3
)

var (
	re     = regexp.MustCompile("[.,!?:;']+")
	cmdMap map[string]func(string) string
)

func init() {
	cmdMap = make(map[string]func(string) string)
	cmdMap["(bin)"] = bin
	cmdMap["(hex)"] = hex
	cmdMap["(up)"] = strings.ToUpper
	cmdMap["(low)"] = strings.ToLower
	cmdMap["(cap)"] = capitalize
	cmdMap["(up"] = strings.ToUpper
	cmdMap["(low"] = strings.ToLower
	cmdMap["(cap"] = capitalize
}

func capitalize(str string) string {
	out := []rune(str)
	r := out[0]
	if r >= 'a' && r <= 'z' {
		out[0] -= alphaOffset
	}
	return string(out)
}

func hex(str string) string {
	n, err := strconv.ParseInt(str, 16, 64)
	if err != nil {
		log.Fatal(err.Error())
	}
	return strconv.Itoa(int(n))
}

func bin(str string) string {
	n, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		log.Fatal(err.Error())
	}
	return strconv.Itoa(int(n))
}
