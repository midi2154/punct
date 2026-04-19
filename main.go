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
