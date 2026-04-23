package transformations

import (
	"regexp"
)


func Quote(s string)string{
	re := regexp.MustCompile(`'\s+(.*?)\s+'`)
	s = re.ReplaceAllString(s, "'$1'")

	re2 := regexp.MustCompile(`"\s+(.*?)\s+"`)
	s = re2.ReplaceAllString(s,`"1$"`)


	re3 := regexp.MustCompile((`(\w+)(\s+)`))
	s = re3.ReplaceAllString(s, `$1 `)


	return s
}


package transformations

import (
	"regexp"
)

func Punct(word string)string{
	Checkpunct := regexp.MustCompile(`\s+([.,:;?!])`)
	word = Checkpunct.ReplaceAllString(word, "$1")

	checkpunct2 := regexp.MustCompile(`([.,:;?!])(\s*)(\w)`)
	word = checkpunct2.ReplaceAllString(word, "$1 $3")

	return word
}

func Nword(s string) string {
	word := strings.Fields(s)

	for i := 0; i < len(word); i++ {
		if (word[i] == "(up," || word[i] == "(low," || word[i] == "(cap,") && i+1 < len(word) {

			word[i+1] = strings.TrimSuffix(word[i+1], ")")
			n, err := strconv.Atoi(word[i+1])
			if err != nil {
				fmt.Println("Error converting")
				return ""
			}

			start := i - n
			if start < 0 {
				start = 0
			}

			for j := start; j < i; j++ {
				if word[i] == "(up," {
					word[j] = strings.ToUpper(word[j])
				}
				if word[i] == "(low," {
					word[j] = strings.ToLower(word[j])
				}
				if word[i] == "(cap," {
					if len(word[j]) > 0 {
						word[j] = strings.ToUpper(string(word[j][0])) + strings.ToLower(word[j][1:])
					}
				}
			}
			word = append(word[:i], word[i+2:]...)
			i--
		}
	}
	return strings.Join(word, " ")
}

func Article(s string) string {
	word := strings.Fields(s)

	for i := 0; i < len(word); i++ {
		if i+1 < len(word) {

			next := strings.Trim(word[i+1], `"'.,!?;:`)

			var ch byte
			for j := 0; j < len(next); j++ {
				if (next[j] >= 'a' && next[j] <= 'z') || (next[j] >= 'A' && next[j] <= 'Z') {
					ch = next[j]
					break
				}
			}

			if ch == 0 {
				continue
			}

			isVowel := strings.ContainsAny(string(ch), "aeiouAEIOU")

			if word[i] == "a" && isVowel {
				word[i] = "an"
			}
			if word[i] == "A" && isVowel {
				word[i] = "An"
			}
			if word[i] == "an" && !isVowel {
				word[i] = "a"
			}
			if word[i] == "An" && !isVowel {
				word[i] = "A"
			}
		}
	}

	return strings.Join(word, " ")
}

