package arrays

import (
	"mlab/rutils"
	"reflect"
	"sort"
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

func TestConcat(t *testing.T) {
	type testData struct {
		arr            []string
		arraysToConcat [][]string
		response       []string
	}

	examples := map[string]testData{
		"empty arr": testData{arr: []string{}, arraysToConcat: [][]string{}, response: []string{}},
		"concat one array": testData{arr: []string{"str1", "str4"}, arraysToConcat: [][]string{[]string{"str2", "str3"}},
			response: []string{"str1", "str4", "str2", "str3"}},
		"concat multiple arrays": testData{arr: []string{"str1", "str4"}, arraysToConcat: [][]string{[]string{"str2", "str3"}, []string{"str5"}},
			response: []string{"str1", "str4", "str2", "str3", "str5"}},
	}

	for k, v := range examples {
		initialArr := StringArray(v.arr)
		initialArr.Concat(v.arraysToConcat...)

		if !reflect.DeepEqual([]string(initialArr), v.response) {
			t.Errorf("test [%v] failed on method Concat with params(initialArr: %v), expected to be %v got %v",
				k, v.arr, v.response, initialArr)
		}
	}
}

func TestIndex(t *testing.T) {
	type testData struct {
		arr      []string
		selector string
		response int
	}

	examples := map[string]testData{
		"empty arr":                  testData{arr: []string{}, selector: "s", response: -1},
		"empty selector":             testData{arr: []string{}, selector: "", response: -1},
		"one occurence in arr":       testData{arr: []string{"str1", "str4"}, selector: "str1", response: 0},
		"multiple occurences in arr": testData{arr: []string{"str1", "str4", "str1"}, selector: "str1", response: 0},
	}

	for k, v := range examples {
		initialArr := StringArray(v.arr)
		resp := initialArr.Index(v.selector)

		if resp != v.response {
			t.Errorf("test [%v] failed on method Index with params(initialArr: %v, selector: %v), expected to be %v got %v",
				k, v.arr, v.selector, v.response, resp)
		}
	}
}

func TestMin(t *testing.T) {
	type testData struct {
		arr      []string
		response string
	}

	examples := map[string]testData{
		"empty arr":                testData{arr: []string{}, response: ""},
		"one element":              testData{arr: []string{"str"}, response: "str"},
		"word with less letters":   testData{arr: []string{"str12", "str123", "str1"}, response: "str1"},
		"words with same length":   testData{arr: []string{"abc", "def", "rty"}, response: "abc"},
		"words with same length 2": testData{arr: []string{"def", "rty", "abc"}, response: "abc"},
	}

	for k, v := range examples {
		initialArr := StringArray(v.arr)
		resp := initialArr.Min()

		if resp != v.response {
			t.Errorf("test [%v] failed on method Min with params(initialArr: %v), expected to be %v got %v",
				k, v.arr, v.response, resp)
		}
	}
}

func TestMax(t *testing.T) {
	type testData struct {
		arr      []string
		response string
	}

	examples := map[string]testData{
		"empty arr":                testData{arr: []string{}, response: ""},
		"one element":              testData{arr: []string{"str"}, response: "str"},
		"word with most letters":   testData{arr: []string{"str12", "str123", "str1"}, response: "str123"},
		"words with same length":   testData{arr: []string{"abc", "def", "rty"}, response: "rty"},
		"words with same length 2": testData{arr: []string{"def", "rty", "abc"}, response: "rty"},
	}

	for k, v := range examples {
		initialArr := StringArray(v.arr)
		resp := initialArr.Max()

		if resp != v.response {
			t.Errorf("test [%v] failed on method Max with params(initialArr: %v), expected to be %v got %v",
				k, v.arr, v.response, resp)
		}
	}
}

func TestPop(t *testing.T) {
	type testData struct {
		arr      []string
		newArr   []string
		n        *int
		response string
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []string{}, newArr: []string{}, response: ""},
		"one element":       testData{arr: []string{"str"}, newArr: []string{}, response: "str"},
		"multiple elements": testData{arr: []string{"str12", "str123", "str1"}, newArr: []string{"str12", "str123"}, response: "str1"},
		"pop two elements":  testData{arr: []string{"str12", "str123", "str1"}, newArr: []string{"str12"}, response: "str123", n: rutils.IntPtr(2)},
		"pop zero elements": testData{arr: []string{"str12", "str123", "str1"}, newArr: []string{"str12", "str123", "str1"}, response: "", n: rutils.IntPtr(0)},
	}

	for k, v := range examples {
		initialArr := StringArray(v.arr)
		var resp string
		if v.n != nil {
			resp = initialArr.Pop(*v.n)
		} else {
			resp = initialArr.Pop()
		}

		if resp != v.response {
			t.Errorf("test [%v] failed on method Pop with params(initialArr: %v), expected to be %v got %v",
				k, v.arr, v.response, resp)
		}

		if !reflect.DeepEqual([]string(initialArr), v.newArr) {
			t.Errorf("test [%v] failed on method Pop with params(initialArr: %v), expected to be %v got %v",
				k, v.arr, initialArr, v.newArr)
		}
	}

}

func TestSelect(t *testing.T) {
	type testData struct {
		arr      []string
		execFunc func(str string) bool
		response []string
	}

	f1 := func(str string) bool {
		if str == "1" {
			return true
		}
		return false
	}

	f2 := func(str string) bool {
		if str == "str123" || str == "str12" {
			return true
		}

		return false
	}

	examples := map[string]testData{
		"empty arr":          testData{arr: []string{}, execFunc: f1, response: []string{}},
		"func returns false": testData{arr: []string{"str"}, execFunc: f1, response: []string{}},
		"valid elements":     testData{arr: []string{"str12", "str123", "str1"}, execFunc: f2, response: []string{"str12", "str123"}},
	}

	for k, v := range examples {
		initialArr := StringArray(v.arr)

		resp := initialArr.Select(v.execFunc)

		if !reflect.DeepEqual(resp, v.response) {
			t.Errorf("test [%v] failed on method Select with params(initialArr: %v), expected to be %v got %v",
				k, v.arr, v.response, resp)
		}

	}
}

func TestUniq(t *testing.T) {
	type testData struct {
		arr      []string
		response []string
	}

	examples := map[string]testData{
		"empty arr":                  testData{arr: []string{}, response: []string{}},
		"just one element":           testData{arr: []string{"str"}, response: []string{"str"}},
		"multiple uniq elements":     testData{arr: []string{"str12", "str123", "str1"}, response: []string{"str12", "str123", "str1"}},
		"multiple non uniq elements": testData{arr: []string{"str12", "str123", "str1", "str12", "str1"}, response: []string{"str12", "str123", "str1"}},
	}

	for k, v := range examples {
		initialArr := StringArray(v.arr)

		initialArr.Uniq()

		sort.Strings([]string(initialArr))
		sort.Strings(v.response)

		if !reflect.DeepEqual([]string(initialArr), v.response) {
			t.Errorf("test [%v] failed on method Uniq with params(initialArr: %v), expected to be %v got %v",
				k, v.arr, v.response, initialArr)
		}

	}

}

func TestToStringArray(t *testing.T) {
	type testData struct {
		arr      []string
		response []string
		err      error
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []string{}, response: []string{}, err: nil},
		"just one element":  testData{arr: []string{"1"}, response: []string{"1"}, err: nil},
		"multiple elements": testData{arr: []string{"2", "3", "1"}, response: []string{"2", "3", "1"}, err: nil},
	}

	for k, v := range examples {
		initialArr := StringArray(v.arr)

		resArr, err := initialArr.ToStringArray()

		if resArr == nil || !reflect.DeepEqual(*resArr, v.response) || err != v.err {
			t.Errorf("test [%v] failed on method ToStringArray with params(initialArr: %v), expected to be %v got %v and error to be %v got %v",
				k, v.arr, v.response, *resArr, v.err, err)
		}

	}
}

func TestToFloat64Array(t *testing.T) {
	type testData struct {
		arr      []string
		response []float64
		err      error
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []string{}, response: []float64{}, err: nil},
		"just one element":  testData{arr: []string{"1"}, response: []float64{1}, err: nil},
		"multiple elements": testData{arr: []string{"2", "3", "1"}, response: []float64{2, 3, 1}, err: nil},
	}

	for k, v := range examples {
		initialArr := StringArray(v.arr)

		resArr, err := initialArr.ToFloat64Array()

		if resArr == nil || !reflect.DeepEqual(*resArr, v.response) || err != v.err {
			t.Errorf("test [%v] failed on method ToFloat64Array with params(initialArr: %v), expected to be %v got %v and error to be %v got %v",
				k, v.arr, v.response, *resArr, v.err, err)
		}

	}
}

func TestToFloat32Array(t *testing.T) {
	type testData struct {
		arr      []string
		response []float32
		err      error
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []string{}, response: []float32{}, err: nil},
		"just one element":  testData{arr: []string{"1"}, response: []float32{1}, err: nil},
		"multiple elements": testData{arr: []string{"2", "3", "1"}, response: []float32{2, 3, 1}, err: nil},
	}

	for k, v := range examples {
		initialArr := StringArray(v.arr)

		resArr, err := initialArr.ToFloat32Array()

		if resArr == nil || !reflect.DeepEqual(*resArr, v.response) || err != v.err {
			t.Errorf("test [%v] failed on method ToFloat32Array with params(initialArr: %v), expected to be %v got %v and error to be %v got %v",
				k, v.arr, v.response, *resArr, v.err, err)
		}

	}
}

func TestToInt64Array(t *testing.T) {
	type testData struct {
		arr      []string
		response []int64
		err      error
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []string{}, response: []int64{}, err: nil},
		"just one element":  testData{arr: []string{"1"}, response: []int64{1}, err: nil},
		"multiple elements": testData{arr: []string{"2", "3", "1"}, response: []int64{2, 3, 1}, err: nil},
	}

	for k, v := range examples {
		initialArr := StringArray(v.arr)

		resArr, err := initialArr.ToInt64Array()

		if resArr == nil || !reflect.DeepEqual(*resArr, v.response) || err != v.err {
			t.Errorf("test [%v] failed on method ToInt64Array with params(initialArr: %v), expected to be %v got %v and error to be %v got %v",
				k, v.arr, v.response, *resArr, v.err, err)
		}
	}
}

func TestToInt32Array(t *testing.T) {
	type testData struct {
		arr      []string
		response []int32
		err      error
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []string{}, response: []int32{}, err: nil},
		"just one element":  testData{arr: []string{"1"}, response: []int32{1}, err: nil},
		"multiple elements": testData{arr: []string{"2", "3", "1"}, response: []int32{2, 3, 1}, err: nil},
	}

	for k, v := range examples {
		initialArr := StringArray(v.arr)

		resArr, err := initialArr.ToInt32Array()

		if resArr == nil || !reflect.DeepEqual(*resArr, v.response) || err != v.err {
			t.Errorf("test [%v] failed on method ToInt32Array with params(initialArr: %v), expected to be %v got %v and error to be %v got %v",
				k, v.arr, v.response, *resArr, v.err, err)
		}
	}
}

func TestToUintArray(t *testing.T) {
	type testData struct {
		arr      []string
		response []uint
		err      error
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []string{}, response: []uint{}, err: nil},
		"just one element":  testData{arr: []string{"1"}, response: []uint{1}, err: nil},
		"multiple elements": testData{arr: []string{"2", "3", "1"}, response: []uint{2, 3, 1}, err: nil},
	}

	for k, v := range examples {
		initialArr := StringArray(v.arr)

		resArr, err := initialArr.ToUintArray()

		if resArr == nil || !reflect.DeepEqual(*resArr, v.response) || err != v.err {
			t.Errorf("test [%v] failed on method ToUintArray with params(initialArr: %v), expected to be %v got %v and error to be %v got %v",
				k, v.arr, v.response, *resArr, v.err, err)
		}
	}
}

func TestToUint32Array(t *testing.T) {
	type testData struct {
		arr      []string
		response []uint32
		err      error
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []string{}, response: []uint32{}, err: nil},
		"just one element":  testData{arr: []string{"1"}, response: []uint32{1}, err: nil},
		"multiple elements": testData{arr: []string{"2", "3", "1"}, response: []uint32{2, 3, 1}, err: nil},
	}

	for k, v := range examples {
		initialArr := StringArray(v.arr)

		resArr, err := initialArr.ToUint32Array()

		if resArr == nil || !reflect.DeepEqual(*resArr, v.response) || err != v.err {
			t.Errorf("test [%v] failed on method ToUint32Array with params(initialArr: %v), expected to be %v got %v and error to be %v got %v",
				k, v.arr, v.response, *resArr, v.err, err)
		}
	}
}

func TestToUint64Array(t *testing.T) {
	type testData struct {
		arr      []string
		response []uint64
		err      error
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []string{}, response: []uint64{}, err: nil},
		"just one element":  testData{arr: []string{"1"}, response: []uint64{1}, err: nil},
		"multiple elements": testData{arr: []string{"2", "3", "1"}, response: []uint64{2, 3, 1}, err: nil},
	}

	for k, v := range examples {
		initialArr := StringArray(v.arr)

		resArr, err := initialArr.ToUint64Array()

		if resArr == nil || !reflect.DeepEqual(*resArr, v.response) || err != v.err {
			t.Errorf("test [%v] failed on method ToUint64Array with params(initialArr: %v), expected to be %v got %v and error to be %v got %v",
				k, v.arr, v.response, *resArr, v.err, err)
		}
	}
}

func TestToJson(t *testing.T) {
	arr := []string{"2", "3", "4"}
	response := `["2","3","4"]`
	var errResp error

	initialArr := StringArray(arr)

	jsonStr, err := initialArr.ToJSON()

	if jsonStr != response || err != errResp {
		t.Errorf("test failed on method ToJSON with params(initialArr: %v), expected to be %v got %v and error to be %v got %v",
			arr, response, jsonStr, errResp, err)
	}

}
