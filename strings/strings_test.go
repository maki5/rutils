package strings

import (
	"testing"

	utils "github.com/maki5/rutils"
)

func TestAt(t *testing.T) {
	type testData struct {
		initialStr string
		selector   interface{}
		response   string
	}
	examples := map[string]testData{
		"first_char":      testData{initialStr: "test_string", selector: 0, response: "t"},
		"first_two_chars": testData{initialStr: "test_string", selector: []int{0, 1}, response: "te"},
	}

	badExamples := map[string]testData{
		"incorrect_params_type":      testData{initialStr: "test_string", selector: "0"},
		"incorrect_number_of_params": testData{initialStr: "test_string", selector: []int{0, 1, 2}},
	}

	for k, v := range examples {
		resp, err := At(v.initialStr, v.selector)
		if resp != v.response {
			t.Errorf("test [%v] failed on method At with params(initialString: %v, selector %v), expected %v got %v (error: %v)",
				k, v.initialStr, v.selector, v.response, resp, err)
		}
	}

	for k, v := range badExamples {
		_, err := At(v.initialStr, v.selector)
		if err == nil {
			t.Errorf("test [%v] failed on method At with params(initialString: %v, selector %v), expected error got %v",
				k, v.initialStr, v.selector, err)
		}
	}
}

func TestBlank(t *testing.T) {
	type testData struct {
		initialStr string
		response   bool
	}

	examples := map[string]testData{
		"empty string":                                 testData{initialStr: "", response: true},
		"empty string with multiple whitespaces":       testData{initialStr: "      ", response: true},
		"non empty string":                             testData{initialStr: "bfhdjsbfhdjs", response: false},
		"non empty string with whitespaces":            testData{initialStr: "     bfhdjsbfhdjs", response: false},
		"non empty string with whitespaces at the end": testData{initialStr: "bfhdjsbfhdjs     ", response: false},
	}

	for k, v := range examples {
		resp := Blank(v.initialStr)
		if v.response != resp {
			t.Errorf("test [%v] failed on method Blank, expected %v got %v", k, v.response, resp)
		}
	}
}

func TestCamelize(t *testing.T) {
	type testData struct {
		initialStr string
		response   string
	}

	examples := map[string]testData{
		"empty string":               testData{initialStr: "", response: ""},
		"string without whitespaces": testData{initialStr: "teststring", response: "Teststring"},
		"only whitespaces string":    testData{initialStr: "   ", response: "   "},
		"string with whitespace":     testData{initialStr: "test string", response: "Test string"},
		"snake case string":          testData{initialStr: "test_string", response: "TestString"},
	}

	for k, v := range examples {
		resp := Camelize(v.initialStr)
		if resp != v.response {
			t.Errorf("test [%v] failed on method Camelize, expected %v got %v", k, v.response, resp)
		}
	}
}

func TestDasherize(t *testing.T) {
	type testData struct {
		initialStr string
		response   string
	}

	examples := map[string]testData{
		"empty string":               testData{initialStr: "", response: ""},
		"string without whitespaces": testData{initialStr: "teststring", response: "teststring"},
		"only whitespaces string":    testData{initialStr: "   ", response: "   "},
		"string with whitespace":     testData{initialStr: "test string", response: "test string"},
		"snake case string":          testData{initialStr: "test_string", response: "test-string"},
	}

	for k, v := range examples {
		resp := Dasherize(v.initialStr)
		if resp != v.response {
			t.Errorf("test [%v] failed on method Dasherize, expected %v got %v", k, v.response, resp)
		}
	}
}

func TestFirst(t *testing.T) {
	type testData struct {
		initialStr string
		selector   *int
		response   string
	}

	examples := map[string]testData{
		"empty string":                       testData{initialStr: "", response: ""},
		"first without selector":             testData{initialStr: "hello", response: "h"},
		"first with selector":                testData{initialStr: "hello", selector: utils.IntPtr(1), response: "h"},
		"first two":                          testData{initialStr: "hello", selector: utils.IntPtr(2), response: "he"},
		"zero":                               testData{initialStr: "hello", selector: utils.IntPtr(0), response: ""},
		"all":                                testData{initialStr: "hello", selector: utils.IntPtr(6), response: "hello"},
		"selector bigger then intial string": testData{initialStr: "hello", selector: utils.IntPtr(7), response: "hello"},
	}

	for k, v := range examples {
		var resp string
		if v.selector == nil {
			resp = First(v.initialStr)
		} else {
			resp = First(v.initialStr, *v.selector)
		}

		if resp != v.response {
			t.Errorf("test [%v] failed on method First, expected %v got %v", k, v.response, resp)
		}
	}
}

func TestFrom(t *testing.T) {
	type testData struct {
		initialStr string
		selector   int
		response   string
	}

	examples := map[string]testData{
		"empty string":              testData{initialStr: "", selector: 1, response: ""},
		"whole word":                testData{initialStr: "hello", selector: 0, response: "hello"},
		"from third":                testData{initialStr: "hello", selector: 3, response: "lo"},
		"negative selector":         testData{initialStr: "hello", selector: -2, response: "lo"},
		"selector too big":          testData{initialStr: "hello", selector: 10, response: ""},
		"negative selector too big": testData{initialStr: "hello", selector: -10, response: "hello"},
	}

	for k, v := range examples {
		var resp string
		resp = From(v.initialStr, v.selector)

		if resp != v.response {
			t.Errorf("test [%v] failed on method From, expected %v got %v", k, v.response, resp)
		}
	}
}

func TestHumanize(t *testing.T) {
	type testData struct {
		initialStr string
		response   string
	}

	examples := map[string]testData{
		"empty string":                 testData{initialStr: "", response: ""},
		"snake case regulat string":    testData{initialStr: "employee_salary", response: "Employee salary"},
		"snake case with id":           testData{initialStr: "author_id", response: "Author"},
		"just `_id`":                   testData{initialStr: "_id", response: "Id"},
		"`_id` with snake case string": testData{initialStr: "_id_test", response: "Test"},
		"non snake case string":        testData{initialStr: "some string", response: "Some string"},
	}

	for k, v := range examples {
		var resp string
		resp = Humanize(v.initialStr)

		if resp != v.response {
			t.Errorf("test [%v] failed on method Humanize, expected %v got %v", k, v.response, resp)
		}
	}
}

func TestLast(t *testing.T) {
	type testData struct {
		initialStr string
		selector   *int
		response   string
	}

	examples := map[string]testData{
		"empty string":                       testData{initialStr: "", response: ""},
		"last without selector":              testData{initialStr: "hello", response: "o"},
		"last with selector":                 testData{initialStr: "hello", selector: utils.IntPtr(1), response: "o"},
		"last two":                           testData{initialStr: "hello", selector: utils.IntPtr(2), response: "lo"},
		"zero":                               testData{initialStr: "hello", selector: utils.IntPtr(0), response: ""},
		"all":                                testData{initialStr: "hello", selector: utils.IntPtr(6), response: "hello"},
		"selector bigger then intial string": testData{initialStr: "hello", selector: utils.IntPtr(7), response: "hello"},
	}

	for k, v := range examples {
		var resp string
		if v.selector == nil {
			resp = Last(v.initialStr)
		} else {
			resp = Last(v.initialStr, *v.selector)
		}

		if resp != v.response {
			t.Errorf("test [%v] failed on method Last, expected %v got %v", k, v.response, resp)
		}
	}
}

// func TestPluralize(t *testing.T) {
// 	type testData struct {
// 		initialStr string
// 		response   string
// 	}

// 	examples := map[string]testData{
// 		"empty str":    testData{initialStr: "", response: ""},
// 		"post":         testData{initialStr: "post", response: "posts"},
// 		"octopus":      testData{initialStr: "octopus", response: "octopi"},
// 		"sheep":        testData{initialStr: "sheep", response: "sheep"},
// 		"words":        testData{initialStr: "words", response: "words"},
// 		"CamelOctopus": testData{initialStr: "CamelOctopus", response: "CamelOctopi"},
// 	}

// 	for k, v := range examples {
// 		var resp string
// 		resp = Pluralize(v.initialStr)

// 		if resp != v.response {
// 			t.Errorf("test [%v] failed on method Pluralize, expected %v got %v", k, v.response, resp)
// 		}
// 	}
// }

func TestSnakeCase(t *testing.T) {
	type testData struct {
		initialStr string
		response   string
	}

	examples := map[string]testData{
		"empty string":           testData{initialStr: "", response: ""},
		"camel case str":         testData{initialStr: "HelloStr", response: "hello_str"},
		"snake case str":         testData{initialStr: "hello_str", response: "hello_str"},
		"simple str":             testData{initialStr: "hello", response: "hello"},
		"complex camel case str": testData{initialStr: "HelloStrStr", response: "hello_str_str"},
		"complex snake case str": testData{initialStr: "hello_str_str", response: "hello_str_str"},
	}

	for k, v := range examples {
		var resp string

		resp = SnakeCase(v.initialStr)

		if resp != v.response {
			t.Errorf("test [%v] failed on method SnakeCase, expected %v got %v", k, v.response, resp)
		}
	}
}

func TestHasOnlyLetters(t *testing.T) {
	type testData struct {
		initialStr string
		response   bool
	}

	examples := map[string]testData{
		"empty string":             testData{initialStr: "", response: false},
		"just letters":             testData{initialStr: "hello", response: true},
		"letters with whitespaces": testData{initialStr: "hello str", response: false},
		"numbers":                  testData{initialStr: "12345", response: false},
		"numbers with letters":     testData{initialStr: "1234 hello", response: false},
	}

	for k, v := range examples {
		resp := HasOnlyLetters(v.initialStr)

		if resp != v.response {
			t.Errorf("test [%v] failed on method HasOnlyLetters, expected %v got %v", k, v.response, resp)
		}
	}
}

func TestInsert(t *testing.T) {
	type testData struct {
		initialStr  string
		index       int
		strToInsert string
		response    string
	}

	examples := map[string]testData{
		"empty string":             testData{initialStr: "", index: 0, strToInsert: "hello", response: "hello"},
		"bad index with empty str": testData{initialStr: "", index: 1, strToInsert: "hello", response: ""},
		"bad index":                testData{initialStr: "hello", index: 6, strToInsert: "dd", response: "hello"},
		"one letter":               testData{initialStr: "hello", index: 2, strToInsert: "l", response: "helllo"},
		"word":                     testData{initialStr: "hello", index: 2, strToInsert: "word", response: "hewordllo"},
		"word with whitespaces":    testData{initialStr: "hello", index: 2, strToInsert: " word ", response: "he word llo"},
		"at start of str":          testData{initialStr: "hello", index: 0, strToInsert: "l", response: "lhello"},
		"at end of str":            testData{initialStr: "hello", index: 5, strToInsert: "l", response: "hellol"},
	}

	for k, v := range examples {
		resp := Insert(v.initialStr, v.index, v.strToInsert)

		if resp != v.response {
			t.Errorf("test [%v] failed on method Insert, expected %v got %v", k, v.response, resp)
		}
	}
}

func TestReverse(t *testing.T) {
	type testData struct {
		initialStr string
		response   string
	}

	examples := map[string]testData{
		"empty string":   testData{initialStr: "", response: ""},
		"regular string": testData{initialStr: "hello", response: "olleh"},
	}

	for k, v := range examples {
		resp := Reverse(v.initialStr)

		if resp != v.response {
			t.Errorf("test [%v] failed on method Reverse, expected %v got %v", k, v.response, resp)
		}
	}
}
