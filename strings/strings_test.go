package strings

import (
	// strings "mlab/rutils/strings"
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
