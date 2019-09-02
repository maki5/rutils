package arrays

import (
	"reflect"
	"testing"
)

func TestDelete(t *testing.T) {
	type testData struct {
		arr      []string
		selector string
		response []string
	}
	examples := map[string]testData{
		"empty arr":                    testData{arr: []string{}, selector: "str1", response: []string{}},
		"sample str arr":               testData{arr: []string{"str1", "str2", "str3", "str4"}, selector: "str1", response: []string{"str2", "str3", "str4"}},
		"sample str arr 1":             testData{arr: []string{"str1", "str2", "str3", "str4"}, selector: "str2", response: []string{"str1", "str3", "str4"}},
		"str doesn't present in array": testData{arr: []string{"str1", "str2", "str3", "str4"}, selector: "str5", response: []string{"str1", "str2", "str3", "str4"}},
	}

	badExamples := map[string]testData{
		"param doesn't exist": testData{arr: []string{"str1", "str2", "str3", "str4"}, selector: "str5", response: []string{"str1", "str2", "str3", "str4", "str5"}},
		"param not deleted":   testData{arr: []string{"str1", "str2", "str3", "str4"}, selector: "str1", response: []string{"str1", "str2", "str3", "str4"}},
	}

	for k, v := range examples {
		initialArr := StringArray(v.arr)

		initialArr.Delete(v.selector)
		if !reflect.DeepEqual([]string(initialArr), v.response) {
			t.Errorf("test [%v] failed on method Delete with params(initialArr: %v, selector %v), expected %v got %v",
				k, v.arr, v.selector, v.response, initialArr)
		}
	}

	for k, v := range badExamples {
		initialArr := StringArray(v.arr)

		initialArr.Delete(v.selector)
		if reflect.DeepEqual([]string(initialArr), v.response) {
			t.Errorf("test [%v] failed on method Delete with params(initialArr: %v, selector %v), expected to not be %v",
				k, v.arr, v.selector, v.response)
		}
	}
}

func TestContains(t *testing.T) {
	type testData struct {
		arr      []string
		selector string
		response bool
	}

	examples := map[string]testData{
		"empty arr":                    testData{arr: []string{}, selector: "str1", response: false},
		"sample str arr":               testData{arr: []string{"str1", "str2", "str3", "str4"}, selector: "str1", response: true},
		"sample str arr 1":             testData{arr: []string{"str1", "str2", "str3", "str4"}, selector: "str2", response: true},
		"str doesn't present in array": testData{arr: []string{"str1", "str2", "str3", "str4"}, selector: "str5", response: false},
	}

	for k, v := range examples {
		arr := StringArray(v.arr)
		resp := arr.Contains(v.selector)
		if v.response != resp {
			t.Errorf("test [%v] failed on method Contains with params(initialArr: %v, selector %v), expected %v got %v",
				k, v.arr, v.selector, v.response, resp)
		}
	}
}

func TestClear(t *testing.T) {
	type testData struct {
		arr []string
	}

	examples := map[string]testData{
		"empty arr":        testData{arr: []string{}},
		"sample str arr":   testData{arr: []string{"str1", "str2", "str3", "str4"}},
		"sample str arr 1": testData{arr: []string{"str1"}},
	}

	for k, v := range examples {
		initialArr := StringArray(v.arr)

		initialArr.Clear()
		arrLen := len(initialArr)
		if arrLen > 0 {
			t.Errorf("test [%v] failed on method Clear with params(initialArr: %v), expected length to be 0 got %v",
				k, v.arr, arrLen)
		}
	}
}

func TestCollect(t *testing.T) {
	type testData struct {
		arr      []string
		execFunc func(el string) string
		response []string
	}

	f1 := func(el string) string {
		return el + "!"
	}

	f2 := func(el string) string {
		if el != "" {
			return el + "!"
		}

		return "empty"
	}

	examples := map[string]testData{
		"empty arr":        testData{arr: []string{}, execFunc: f1, response: []string{}},
		"sample str arr":   testData{arr: []string{"str1", "str2", "str3", "str4"}, execFunc: f1, response: []string{"str1!", "str2!", "str3!", "str4!"}},
		"sample str arr 1": testData{arr: []string{"str1", "", "str2", "", "str3"}, execFunc: f2, response: []string{"str1!", "empty", "str2!", "empty", "str3!"}},
	}

	for k, v := range examples {
		arr := StringArray(v.arr)
		newArr := arr.Collect(v.execFunc)

		if !reflect.DeepEqual(newArr, v.response) {
			t.Errorf("test [%v] failed on method Collect with params(initialArr: %v), expected to be %v got %v",
				k, v.arr, v.response, newArr)
		}
	}
}

func TestCompact(t *testing.T) {
	type testData struct {
		arr      []string
		response []string
	}

	examples := map[string]testData{
		"empty arr":        testData{arr: []string{}, response: []string{}},
		"sample str arr":   testData{arr: []string{"str1", "", "    ", "str4"}, response: []string{"str1", "str4"}},
		"sample str arr 1": testData{arr: []string{"", "str1", "", "str2", "", "str3"}, response: []string{"str1", "str2", "str3"}},
	}

	for k, v := range examples {
		initialArr := StringArray(v.arr)
		initialArr.Compact()

		if !reflect.DeepEqual([]string(initialArr), v.response) {
			t.Errorf("test [%v] failed on method Collect with params(initialArr: %v), expected to be %v got %v",
				k, v.arr, v.response, initialArr)
		}
	}

}
