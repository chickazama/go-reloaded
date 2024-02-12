package main

import (
	"log"
	"os"
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
	cmdMap["(bin"] = bin
	cmdMap["(hex"] = hex
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

func tokenizeInput(str string) []string {
	fields := strings.Fields(str)
	var ret []string
	for _, f := range fields {
		locs := re.FindAllStringIndex(f, -1)
		if locs != nil {
			idxs := make(map[int]bool)
			for _, loc := range locs {
				for _, v := range loc {
					idxs[v] = true
				}
			}
			var sb strings.Builder
			for i, r := range f {
				if idxs[i] {
					if sb.Len() > 0 {
						ret = append(ret, sb.String())
					}
					sb.Reset()
				}
				sb.WriteRune(r)
			}
			ret = append(ret, sb.String())
		} else {
			ret = append(ret, f)
		}
	}
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

func executeCommands(tokens []string) []string {
	var ret []string
	for i := 0; i < len(tokens); i++ {
		w := tokens[i]
		target := strings.TrimSuffix(w, ")")
		f := cmdMap[target]
		switch w {
		case "(up)", "(low)", "(cap)", "(bin)", "(hex)":
			count := 0
			j := 1
			for count < 1 && len(ret)-j >= 0 {
				if !re.MatchString(ret[len(ret)-j]) {
					ret[len(ret)-j] = f(ret[len(ret)-j])
					count++
				}
				j++
			}
		case "(up", "(cap", "(low":
			if i+2 < len(tokens) {
				n := getN(tokens[i+2])
				count := 0
				j := 1
				for count < n && len(ret)-j >= 0 {
					if !re.MatchString(ret[len(ret)-j]) {
						ret[len(ret)-j] = f(ret[len(ret)-j])
						count++
					}
					j++
				}
				i += 2
			}
		case "a", "an":
			if i+1 >= len(tokens) {
				ret = append(ret, w)
			} else {
				next := strings.ToLower(string(tokens[i+1][0]))
				switch next {
				case "a", "e", "i", "o", "u", "h":
					ret = append(ret, "an")
				default:
					ret = append(ret, "a")
				}
			}
		default:
			ret = append(ret, w)
		}
	}
	return ret
}

func buildOutput(tokens []string) string {
	var sb strings.Builder
	q := false
	for i := range tokens {
		sb.WriteString(tokens[i])
		if tokens[i] == "'" {
			q = !q
			if q {
				continue
			}
		}
		if i < len(tokens)-1 && !re.MatchString(tokens[i+1]) {
			sb.WriteRune(' ')
		} else if i < len(tokens)-1 && tokens[i+1] == "'" && !q {
			sb.WriteRune(' ')
		}
	}
	return sb.String()
}

func main() {
	if len(os.Args) != expArgc {
		log.Fatal("invalid argument count. expect 2 arguments corresponding to input/output file names")
	}
	inpath := os.Args[1]
	outpath := os.Args[2]
	buf, err := os.ReadFile(inpath)
	if err != nil {
		log.Fatal(err.Error())
	}
	tokens := tokenizeInput(string(buf))
	exe := executeCommands(tokens)
	output := buildOutput(exe)
	err = os.WriteFile(outpath, []byte(output), 0664)
	if err != nil {
		log.Fatal(err.Error())
	}
}
