package main

import "regexp"

const (
	alphaOffset = 'a' - 'A'
	expArgc     = 3
)

var (
	re     = regexp.MustCompile("[.,!?:;']+")
	cmdMap map[string]func(string) string
)
