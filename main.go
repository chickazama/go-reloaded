package main

const (
	alphaOffset = 'a' - 'A'
	expArgc     = 3
)

// var (
// 	re     = regexp.MustCompile("[.,!?:;']+")
// 	cmdMap map[string]func(string) string
// )

// func init() {
// 	cmdMap = make(map[string]func(string) string)
// 	cmdMap["(bin)"] = bin
// 	cmdMap["(hex)"] = hex
// 	cmdMap["(up)"] = strings.ToUpper
// 	cmdMap["(low)"] = strings.ToLower
// 	cmdMap["(cap)"] = capitalize
// 	cmdMap["(up"] = strings.ToUpper
// 	cmdMap["(low"] = strings.ToLower
// 	cmdMap["(cap"] = capitalize
// }

// func capitalize(str string) string {
// 	out := []rune(str)
// 	r := out[0]
// 	if r >= 'a' && r <= 'z' {
// 		out[0] -= alphaOffset
// 	}
// 	return string(out)
// }

// func hex(str string) string {
// 	n, err := strconv.ParseInt(str, 16, 64)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	return strconv.Itoa(int(n))
// }

// func bin(str string) string {
// 	n, err := strconv.ParseInt(str, 2, 64)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	return strconv.Itoa(int(n))
// }

// func getN(str string) int {
// 	var runes []rune
// 	for _, r := range str {
// 		if r >= '0' && r <= '9' {
// 			runes = append(runes, r)
// 		}
// 	}
// 	if len(runes) <= 0 {
// 		log.Fatal("not valid command")
// 	}
// 	n, err := strconv.Atoi(string(runes))
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	return n
// }

// func tokenize(str string) []string {
// 	// Separate out all fields of original input
// 	fields := strings.Fields(str)
// 	// fmt.Println("FIELDS")
// 	// for i, val := range fields {
// 	// 	fmt.Printf("[%d]: %s\n", i, val)
// 	// }
// 	// fmt.Println()
// 	var ret []string
// 	// For each field, separate out all punctuation tokens
// 	for _, f := range fields {
// 		// Get the indices of occurrences of punctation expressions
// 		locs := re.FindAllStringIndex(f, -1)
// 		if locs != nil {
// 			// Create a map to store reference to punctuation indices
// 			idxs := make(map[int]bool)
// 			for _, loc := range locs {
// 				for _, v := range loc {
// 					idxs[v] = true
// 				}
// 			}
// 			// Define a string builder to build next token
// 			var sb strings.Builder
// 			// Iterate through each rune in the field
// 			// If the index of that rune is in the map,
// 			// append non-empty strings to the token set
// 			// and reset the string builder.
// 			for i, r := range f {
// 				if idxs[i] {
// 					if sb.Len() > 0 {
// 						ret = append(ret, sb.String())
// 					}
// 					sb.Reset()
// 				}
// 				// Write rune to string
// 				sb.WriteRune(r)
// 			}
// 			// After all runes have been examined,
// 			// append remaining string to return token set
// 			ret = append(ret, sb.String())
// 		} else {
// 			// No tokens found, append field 'as is'.
// 			ret = append(ret, f)
// 		}
// 	}
// 	fmt.Println("TOKENS")
// 	for i, val := range ret {
// 		fmt.Printf("[%d]: %s\n", i, val)
// 	}
// 	fmt.Println()
// 	return ret
// }

// func execute(tokens []string) []string {
// 	var ret []string
// 	for i := 0; i < len(tokens); i++ {
// 		w := tokens[i]
// 		f := cmdMap[w]
// 		switch w {
// 		case "(up)", "(low)", "(cap)", "(bin)", "(hex)":
// 			count := 0
// 			j := 1
// 			for count < 1 && len(ret)-j >= 0 {
// 				if !re.MatchString(ret[len(ret)-j]) {
// 					ret[len(ret)-j] = f(ret[len(ret)-j])
// 					count++
// 				}
// 				j++
// 			}
// 		case "(up", "(cap", "(low":
// 			if i+2 < len(tokens) {
// 				n := getN(tokens[i+2])
// 				count := 0
// 				j := 1
// 				for count < n && len(ret)-j >= 0 {
// 					if !re.MatchString(ret[len(ret)-j]) {
// 						ret[len(ret)-j] = f(ret[len(ret)-j])
// 						count++
// 					}
// 					j++
// 				}
// 				i += 2
// 			}
// 		default:
// 			ret = append(ret, w)
// 		}
// 	}
// 	fmt.Println("EXECUTED")
// 	for i, val := range ret {
// 		fmt.Printf("[%d]: %s\n", i, val)
// 	}
// 	fmt.Println()
// 	return ret
// }

// func main() {
// 	// Get Command-Line Arguments
// 	argc := len(os.Args)
// 	if argc != expArgc {
// 		msg := fmt.Sprintf("invalid argument count. expect 2 arguments corresponding to input/output file names. actual: %d\n", argc-1)
// 		log.Fatal(msg)
// 	}
// 	inpath := os.Args[1]
// 	outpath := os.Args[2]
// 	// Read input file into buffer
// 	buf, err := os.ReadFile(inpath)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	tokens := tokenize(string(buf))
// 	exe := execute(tokens)
// 	b := build(exe)
// 	fmt.Println(b)
// 	err = os.WriteFile(outpath, []byte(b), 0664)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// }

// func build(tokens []string) string {
// 	var sb strings.Builder
// 	q := false
// 	for i := range tokens {
// 		sb.WriteString(tokens[i])
// 		if tokens[i] == "'" {
// 			q = !q
// 			if q {
// 				continue
// 			}
// 		}
// 		if i < len(tokens)-1 && !re.MatchString(tokens[i+1]) {
// 			sb.WriteRune(' ')
// 		} else if i < len(tokens)-1 && tokens[i+1] == "'" && !q {
// 			sb.WriteRune(' ')
// 		}
// 	}
// 	return sb.String()
// }
