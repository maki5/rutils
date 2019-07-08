package strings

import (
	"fmt"
	"mlab/rutils"
	arrayOfStrings "mlab/rutils/arrays/of_string"
	strings2 "strings"
	"unicode"
)

// At : Returns the substring of provided position
//
//
// str := "test_string"
//
// At(0)      # => "t"
//
// At([]int{0, 1})   # => "te"
func At(str string, pos interface{}) (string, error) {
	if p, ok := pos.(int); ok {
		return string(str[p]), nil
	}

	if p, ok := pos.([]int); ok {
		if len(p) != 2 {
			return "", fmt.Errorf("wrong params, expected 2 got %v", len(p))
		}

		return string(str[p[0] : p[1]+1]), nil
	}
	return "", fmt.Errorf("wrong params")
}

// Blank :
func Blank(str string) bool {
	if len(str) == 0 {
		return true
	}

	notBlank := false

	for _, s := range str {
		if string(s) != " " {
			notBlank = true
			break
		}
	}

	if notBlank {
		return false
	} else {
		return true
	}
}

func Camelize(str string) string {
	stringsArr := strings2.Split(str, "_")

	if len(stringsArr) == 0 {
		return str
	}

	var result string
	for _, e := range stringsArr {
		if len(e) == 0 {
			continue
		}

		firstLetter := e[0]
		remainingChars := e[1:]

		result += strings2.ToUpper(string(firstLetter))
		result += remainingChars
	}

	return result
}

func Capitalize(str string) string {
	if len(str) == 0 {
		return ""
	}
	firstLetter := str[0]
	remainingChars := str[1:]

	var result string
	result += strings2.ToUpper(string(firstLetter))
	result += remainingChars

	return result
}

func Dasherize(str string) string {
	return strings2.Replace(str, "_", "-", -1)
}

func First(str string, selectorArgs ...int) string {
	if len(str) == 0 {
		return ""
	}

	selector := 0
	if len(selectorArgs) > 0 {
		selector = selectorArgs[0]
	} else {
		return string(str[0])
	}

	if selector == 0 {
		return ""
	} else {
		if selector > len(str) {
			return str
		}
		return str[:selector]
	}
}

func From(str string, selector int) string {
	if len(str) == 0 || selector > len(str) {
		return ""
	}

	if selector < 0 {
		positiveSelector := rutils.InverseInt(selector)
		if positiveSelector < len(str) {
			validSelector := len(str) - positiveSelector
			return str[validSelector:]
		} else {
			return str
		}
	}

	return str[selector:]
}

func Humanize(str string) string {
	if len(str) == 0 {
		return ""
	}

	var result string
	if str == "_id" {
		return "Id"
	}

	str = strings2.Replace(str, "_id", "", -1)
	stringsArr := strings2.Split(str, "_")
	arrayOfStrings.DeleteString(&stringsArr, "")

	for i, e := range stringsArr {
		if i == 0 {
			result += Capitalize(e)
		} else {
			result += e
		}

		if i < len(stringsArr)-1 {
			result += " "
		}
	}

	return result
}

func Last(str string, selectorArgs ...int) string {
	if len(str) == 0 {
		return ""
	}

	selector := 0
	if len(selectorArgs) > 0 {
		selector = selectorArgs[0]
	} else {
		return string(str[len(str)-1])
	}

	if selector == 0 {
		return ""
	} else {
		if selector > len(str) {
			return str
		}
		selector = len(str) - selector
		return str[selector:]
	}
}

// func Pluralize(str string) string {
// 	if len(str) == 0 {
// 		return str
// 	}

// 	i := inflections{}
// 	uncountable := i.uncountable()
// 	if OfStrings.Contains(&uncountable, str) {
// 		return str
// 	}

// 	for _, e := range i.plural() {
// 		var re = regexp.MustCompile(e.rule)
// 		s := re.ReplaceAllString(str, e.replacement)
// 		if s != str {
// 			return s
// 		}
// 	}

// 	return str
// }

func SnakeCase(str string) string {
	wordsArr := make([]string, 0)
	word := make([]rune, 0)
	var result string

	for i, e := range str {
		if i != 0 && unicode.IsUpper(e) {
			wordsArr = append(wordsArr, string(word))
			word = word[:0]
		}

		runes := []rune(strings2.ToLower(string(e)))
		word = append(word, runes[0])

		if i == len(str)-1 {
			wordsArr = append(wordsArr, string(word))
			word = word[:0]
		}
	}

	for i, e := range wordsArr {
		result += e
		if i < len(wordsArr)-1 {
			result += "_"
		}
	}
	return result
}

func HasOnlyLetters(str string) bool {
	if Blank(str) {
		return false
	}

	for _, r := range str {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func Insert(str string, index int, strToInsert string) string {
	if Blank(str) {
		if len(str) < index {
			return str
		}

		return strToInsert
	}

	if len(str) < index {
		return str
	}

	firstPart := str[:index]
	secondPart := str[index:]

	result := firstPart + strToInsert + secondPart
	return result
}

func Reverse(str string) string {
	if Blank(str) {
		return str
	}

	var result string
	for i := len(str); i > 0; i-- {
		result += string(str[i-1])
	}
	return result
}
