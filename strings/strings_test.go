package strings

import (
	"testing"
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
			t.Errorf("test [%v] failed with params(initialString: %v, selector %v), expected %v got %v (error: %v)",
				k, v.initialStr, v.selector, v.response, resp, err)
		}
	}

	for k, v := range badExamples {
		_, err := At(v.initialStr, v.selector)
		if err == nil {
			t.Errorf("test [%v] failed with params(initialString: %v, selector %v), expected error got %v",
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
		"empty string":                           testData{initialStr: "", response: true},
		"empty string with multiple whitespaces": testData{initialStr: "      ", response: true},
		"non empty string": testData{initialStr: "bfhdjsbfhdjs", response: false},
		"non empty string with whitespaces": testData{initialStr: "     bfhdjsbfhdjs", response: false},
		"non empty string with whitespaces at the end": testData{initialStr: "bfhdjsbfhdjs     ", response: false},
	}

	for k, v := range examples {
		resp := Blank(v.initialStr)
		if v.response != resp {
			t.Errorf("test [%v] failed, expected %v got %v", k, v.response, resp)
		}
	}
}
