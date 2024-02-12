package main

import (
	"fmt"
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

func tokenize(str string) []string {
	// Separate out all fields of original input
	fields := strings.Fields(str)
	// fmt.Println("FIELDS")
	// for i, val := range fields {
	// 	fmt.Printf("[%d]: %s\n", i, val)
	// }
	// fmt.Println()
	var ret []string
	// For each field, separate out all punctuation tokens
	for _, f := range fields {
		// Get the indices of occurrences of punctation expressions
		locs := re.FindAllStringIndex(f, -1)
		if locs != nil {
			// Create a map to store reference to punctuation indices
			idxs := make(map[int]bool)
			for _, loc := range locs {
				for _, v := range loc {
					idxs[v] = true
				}
			}
			// Define a string builder to build next token
			var sb strings.Builder
			// Iterate through each rune in the field
			// If the index of that rune is in the map,
			// append non-empty strings to the token set
			// and reset the string builder.
			for i, r := range f {
				if idxs[i] {
					if sb.Len() > 0 {
						ret = append(ret, sb.String())
					}
					sb.Reset()
				}
				// Write rune to string
				sb.WriteRune(r)
			}
			// After all runes have been examined,
			// append remaining string to return token set
			ret = append(ret, sb.String())
		} else {
			// No tokens found, append field 'as is'.
			ret = append(ret, f)
		}
	}
	fmt.Println("TOKENS")
	for i, val := range ret {
		fmt.Printf("[%d]: %s\n", i, val)
	}
	fmt.Println()
	return ret
}

func getN(str string) int {
	var runes []rune
	for _, r := range str {
		if r >= '0' && r <= '9' {
			runes = append(runes, r)
		}
	}
	if len(runes) <= 0 {
		log.Fatal("not valid command")
	}
	n, err := strconv.Atoi(string(runes))
	if err != nil {
		log.Fatal(err.Error())
	}
	return n
}
