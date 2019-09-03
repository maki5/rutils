package arrays

import (
	"reflect"
	"sort"
	"testing"

	"github.com/maki5/rutils"
)

func TestDelete(t *testing.T) {
	type testData struct {
		arr      []int
		selector int
		response []int
	}
	examples := map[string]testData{
		"empty arr":                    testData{arr: []int{}, selector: 1, response: []int{}},
		"sample int arr":               testData{arr: []int{1, 2, 3, 4}, selector: 1, response: []int{2, 3, 4}},
		"sample int arr 1":             testData{arr: []int{1, 2, 3, 4}, selector: 2, response: []int{1, 3, 4}},
		"int doesn't present in array": testData{arr: []int{1, 2, 3, 4}, selector: 5, response: []int{1, 2, 3, 4}},
	}

	badExamples := map[string]testData{
		"param doesn't exist": testData{arr: []int{1, 2, 3, 4}, selector: 5, response: []int{1, 2, 3, 4, 5}},
		"param not deleted":   testData{arr: []int{1, 2, 3, 4}, selector: 1, response: []int{1, 2, 3, 4}},
	}

	for k, v := range examples {
		initialArr := IntArray(v.arr)

		initialArr.Delete(v.selector)
		if !reflect.DeepEqual([]int(initialArr), v.response) {
			t.Errorf("test [%v] failed on method Delete with params(initialArr: %v, selector %v), expected %v got %v",
				k, v.arr, v.selector, v.response, initialArr)
		}
	}

	for k, v := range badExamples {
		initialArr := IntArray(v.arr)

		initialArr.Delete(v.selector)
		if reflect.DeepEqual([]int(initialArr), v.response) {
			t.Errorf("test [%v] failed on method Delete with params(initialArr: %v, selector %v), expected to not be %v",
				k, v.arr, v.selector, v.response)
		}
	}
}

func TestContains(t *testing.T) {
	type testData struct {
		arr      []int
		selector int
		response bool
	}

	examples := map[string]testData{
		"empty arr":                    testData{arr: []int{}, selector: 1, response: false},
		"sample str arr":               testData{arr: []int{1, 2, 3, 4}, selector: 1, response: true},
		"sample str arr 1":             testData{arr: []int{1, 2, 3, 4}, selector: 2, response: true},
		"str doesn't present in array": testData{arr: []int{1, 2, 3, 4}, selector: 5, response: false},
	}

	for k, v := range examples {
		arr := IntArray(v.arr)
		resp := arr.Contains(v.selector)
		if v.response != resp {
			t.Errorf("test [%v] failed on method Contains with params(initialArr: %v, selector %v), expected %v got %v",
				k, v.arr, v.selector, v.response, resp)
		}
	}
}

func TestClear(t *testing.T) {
	type testData struct {
		arr []int
	}

	examples := map[string]testData{
		"empty arr":        testData{arr: []int{}},
		"sample str arr":   testData{arr: []int{1, 2, 3, 4}},
		"sample str arr 1": testData{arr: []int{1}},
	}

	for k, v := range examples {
		initialArr := IntArray(v.arr)

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
		arr      []int
		execFunc func(el int) int
		response []int
	}

	f1 := func(el int) int {
		return el + 1
	}

	f2 := func(el int) int {
		if el != 0 {
			return el + 1
		}

		return -1
	}

	examples := map[string]testData{
		"empty arr":        testData{arr: []int{}, execFunc: f1, response: []int{}},
		"sample str arr":   testData{arr: []int{1, 2, 3, 4}, execFunc: f1, response: []int{2, 3, 4, 5}},
		"sample str arr 1": testData{arr: []int{1, 0, 2, 0, 3}, execFunc: f2, response: []int{2, -1, 3, -1, 4}},
	}

	for k, v := range examples {
		arr := IntArray(v.arr)
		newArr := arr.Collect(v.execFunc)

		if !reflect.DeepEqual(newArr, v.response) {
			t.Errorf("test [%v] failed on method Collect with params(initialArr: %v), expected to be %v got %v",
				k, v.arr, v.response, newArr)
		}
	}
}

func TestConcat(t *testing.T) {
	type testData struct {
		arr            []int
		arraysToConcat [][]int
		response       []int
	}

	examples := map[string]testData{
		"empty arr": testData{arr: []int{}, arraysToConcat: [][]int{}, response: []int{}},
		"concat one array": testData{arr: []int{1, 4}, arraysToConcat: [][]int{[]int{2, 3}},
			response: []int{1, 4, 2, 3}},
		"concat multiple arrays": testData{arr: []int{1, 4}, arraysToConcat: [][]int{[]int{2, 3}, []int{5}},
			response: []int{1, 4, 2, 3, 5}},
	}

	for k, v := range examples {
		initialArr := IntArray(v.arr)
		initialArr.Concat(v.arraysToConcat...)

		if !reflect.DeepEqual([]int(initialArr), v.response) {
			t.Errorf("test [%v] failed on method Concat with params(initialArr: %v), expected to be %v got %v",
				k, v.arr, v.response, initialArr)
		}
	}
}

func TestIndex(t *testing.T) {
	type testData struct {
		arr      []int
		selector int
		response *int
	}

	examples := map[string]testData{
		"empty arr":                  testData{arr: []int{}, selector: 1, response: nil},
		"one occurence in arr":       testData{arr: []int{1, 4}, selector: 1, response: rutils.IntPtr(0)},
		"multiple occurences in arr": testData{arr: []int{1, 4, 1}, selector: 1, response: rutils.IntPtr(0)},
	}

	for k, v := range examples {
		initialArr := IntArray(v.arr)
		resp := initialArr.Index(v.selector)

		if resp != v.response {
			if resp != nil && v.response != nil && *resp != *v.response {
				t.Errorf("test [%v] failed on method Index with params(initialArr: %v, selector: %v), expected to be %v got %v",
					k, v.arr, v.selector, v.response, resp)
			}
		}
	}
}

func TestMin(t *testing.T) {
	type testData struct {
		arr      []int
		response *int
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []int{}, response: nil},
		"one element":       testData{arr: []int{1}, response: rutils.IntPtr(1)},
		"multiple elements": testData{arr: []int{1, 3, 4, 7, 3, 2}, response: rutils.IntPtr(1)},
	}

	for k, v := range examples {
		initialArr := IntArray(v.arr)
		resp := initialArr.Min()

		if resp != v.response {
			if resp != nil && v.response != nil && *resp != *v.response {
				t.Errorf("test [%v] failed on method Min with params(initialArr: %v), expected to be %v got %v",
					k, v.arr, v.response, resp)
			}
		}
	}
}

func TestMax(t *testing.T) {
	type testData struct {
		arr      []int
		response *int
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []int{}, response: nil},
		"one element":       testData{arr: []int{1}, response: rutils.IntPtr(1)},
		"multiple elements": testData{arr: []int{1, 3, 4, 7, 3, 2}, response: rutils.IntPtr(7)},
	}

	for k, v := range examples {
		initialArr := IntArray(v.arr)
		resp := initialArr.Max()

		if resp != v.response {
			if resp != nil && v.response != nil && *resp != *v.response {
				t.Errorf("test [%v] failed on method Max with params(initialArr: %v), expected to be %v got %v",
					k, v.arr, v.response, resp)
			}
		}
	}
}

func TestPop(t *testing.T) {
	type testData struct {
		arr      []int
		newArr   []int
		n        *int
		response *int
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []int{}, newArr: []int{}, response: nil},
		"one element":       testData{arr: []int{1}, newArr: []int{}, response: rutils.IntPtr(1)},
		"multiple elements": testData{arr: []int{1, 2, 3}, newArr: []int{1, 2}, response: rutils.IntPtr(3)},
		"pop two elements":  testData{arr: []int{1, 2, 3}, newArr: []int{1}, response: rutils.IntPtr(2), n: rutils.IntPtr(2)},
		"pop zero elements": testData{arr: []int{1, 2, 3}, newArr: []int{1, 2, 3}, response: nil, n: rutils.IntPtr(0)},
	}

	for k, v := range examples {
		initialArr := IntArray(v.arr)
		var resp *int
		if v.n != nil {
			resp = initialArr.Pop(*v.n)
		} else {
			resp = initialArr.Pop()
		}

		if resp != v.response {
			if resp != nil && v.response != nil && *resp != *v.response {
				t.Errorf("test [%v] failed on method Pop with params(initialArr: %v), expected to be %v got %v",
					k, v.arr, v.response, resp)
			}
		}

		if !reflect.DeepEqual([]int(initialArr), v.newArr) {
			t.Errorf("test [%v] failed on method Pop with params(initialArr: %v), expected to be %v got %v",
				k, v.arr, initialArr, v.newArr)
		}
	}

}

func TestSelect(t *testing.T) {
	type testData struct {
		arr      []int
		execFunc func(elem int) bool
		response []int
	}

	f1 := func(elem int) bool {
		if elem == 1 {
			return true
		}
		return false
	}

	f2 := func(elem int) bool {
		if elem == 3 || elem == 2 {
			return true
		}

		return false
	}

	examples := map[string]testData{
		"empty arr":          testData{arr: []int{}, execFunc: f1, response: []int{}},
		"func returns false": testData{arr: []int{2}, execFunc: f1, response: []int{}},
		"valid elements":     testData{arr: []int{2, 3, 1}, execFunc: f2, response: []int{2, 3}},
	}

	for k, v := range examples {
		initialArr := IntArray(v.arr)

		resp := initialArr.Select(v.execFunc)

		if !reflect.DeepEqual(resp, v.response) {
			t.Errorf("test [%v] failed on method Select with params(initialArr: %v), expected to be %v got %v",
				k, v.arr, v.response, resp)
		}

	}
}

func TestUniq(t *testing.T) {
	type testData struct {
		arr      []int
		response []int
	}

	examples := map[string]testData{
		"empty arr":                  testData{arr: []int{}, response: []int{}},
		"just one element":           testData{arr: []int{1}, response: []int{1}},
		"multiple uniq elements":     testData{arr: []int{2, 3, 1}, response: []int{2, 3, 1}},
		"multiple non uniq elements": testData{arr: []int{2, 3, 1, 2, 1}, response: []int{2, 3, 1}},
	}

	for k, v := range examples {
		initialArr := IntArray(v.arr)

		initialArr.Uniq()

		sort.Ints([]int(initialArr))
		sort.Ints(v.response)

		if !reflect.DeepEqual([]int(initialArr), v.response) {
			t.Errorf("test [%v] failed on method Uniq with params(initialArr: %v), expected to be %v got %v",
				k, v.arr, v.response, initialArr)
		}

	}

}

func TestToStringArray(t *testing.T) {
	type testData struct {
		arr      []int
		response []string
		err      error
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []int{}, response: []string{}, err: nil},
		"just one element":  testData{arr: []int{1}, response: []string{"1"}, err: nil},
		"multiple elements": testData{arr: []int{2, 3, 1}, response: []string{"2", "3", "1"}, err: nil},
	}

	for k, v := range examples {
		initialArr := IntArray(v.arr)

		resArr, err := initialArr.ToStringArray()

		if resArr == nil || !reflect.DeepEqual(*resArr, v.response) || err != v.err {
			t.Errorf("test [%v] failed on method ToStringArray with params(initialArr: %v), expected to be %v got %v and error to be %v got %v",
				k, v.arr, v.response, *resArr, v.err, err)
		}

	}
}

func TestToFloat64Array(t *testing.T) {
	type testData struct {
		arr      []int
		response []float64
		err      error
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []int{}, response: []float64{}, err: nil},
		"just one element":  testData{arr: []int{1}, response: []float64{1}, err: nil},
		"multiple elements": testData{arr: []int{2, 3, 1}, response: []float64{2, 3, 1}, err: nil},
	}

	for k, v := range examples {
		initialArr := IntArray(v.arr)

		resArr, err := initialArr.ToFloat64Array()

		if resArr == nil || !reflect.DeepEqual(*resArr, v.response) || err != v.err {
			t.Errorf("test [%v] failed on method ToFloat64Array with params(initialArr: %v), expected to be %v got %v and error to be %v got %v",
				k, v.arr, v.response, *resArr, v.err, err)
		}

	}
}

func TestToFloat32Array(t *testing.T) {
	type testData struct {
		arr      []int
		response []float32
		err      error
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []int{}, response: []float32{}, err: nil},
		"just one element":  testData{arr: []int{1}, response: []float32{1}, err: nil},
		"multiple elements": testData{arr: []int{2, 3, 1}, response: []float32{2, 3, 1}, err: nil},
	}

	for k, v := range examples {
		initialArr := IntArray(v.arr)

		resArr, err := initialArr.ToFloat32Array()

		if resArr == nil || !reflect.DeepEqual(*resArr, v.response) || err != v.err {
			t.Errorf("test [%v] failed on method ToFloat32Array with params(initialArr: %v), expected to be %v got %v and error to be %v got %v",
				k, v.arr, v.response, *resArr, v.err, err)
		}

	}
}

func TestToInt64Array(t *testing.T) {
	type testData struct {
		arr      []int
		response []int64
		err      error
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []int{}, response: []int64{}, err: nil},
		"just one element":  testData{arr: []int{1}, response: []int64{1}, err: nil},
		"multiple elements": testData{arr: []int{2, 3, 1}, response: []int64{2, 3, 1}, err: nil},
	}

	for k, v := range examples {
		initialArr := IntArray(v.arr)

		resArr, err := initialArr.ToInt64Array()

		if resArr == nil || !reflect.DeepEqual(*resArr, v.response) || err != v.err {
			t.Errorf("test [%v] failed on method ToInt64Array with params(initialArr: %v), expected to be %v got %v and error to be %v got %v",
				k, v.arr, v.response, *resArr, v.err, err)
		}
	}
}

func TestToInt32Array(t *testing.T) {
	type testData struct {
		arr      []int
		response []int32
		err      error
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []int{}, response: []int32{}, err: nil},
		"just one element":  testData{arr: []int{1}, response: []int32{1}, err: nil},
		"multiple elements": testData{arr: []int{2, 3, 1}, response: []int32{2, 3, 1}, err: nil},
	}

	for k, v := range examples {
		initialArr := IntArray(v.arr)

		resArr, err := initialArr.ToInt32Array()

		if resArr == nil || !reflect.DeepEqual(*resArr, v.response) || err != v.err {
			t.Errorf("test [%v] failed on method ToInt32Array with params(initialArr: %v), expected to be %v got %v and error to be %v got %v",
				k, v.arr, v.response, *resArr, v.err, err)
		}
	}
}

func TestToUintArray(t *testing.T) {
	type testData struct {
		arr      []int
		response []uint
		err      error
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []int{}, response: []uint{}, err: nil},
		"just one element":  testData{arr: []int{1}, response: []uint{1}, err: nil},
		"multiple elements": testData{arr: []int{2, 3, 1}, response: []uint{2, 3, 1}, err: nil},
	}

	for k, v := range examples {
		initialArr := IntArray(v.arr)

		resArr, err := initialArr.ToUintArray()

		if resArr == nil || !reflect.DeepEqual(*resArr, v.response) || err != v.err {
			t.Errorf("test [%v] failed on method ToUintArray with params(initialArr: %v), expected to be %v got %v and error to be %v got %v",
				k, v.arr, v.response, *resArr, v.err, err)
		}
	}
}

func TestToUint32Array(t *testing.T) {
	type testData struct {
		arr      []int
		response []uint32
		err      error
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []int{}, response: []uint32{}, err: nil},
		"just one element":  testData{arr: []int{1}, response: []uint32{1}, err: nil},
		"multiple elements": testData{arr: []int{2, 3, 1}, response: []uint32{2, 3, 1}, err: nil},
	}

	for k, v := range examples {
		initialArr := IntArray(v.arr)

		resArr, err := initialArr.ToUint32Array()

		if resArr == nil || !reflect.DeepEqual(*resArr, v.response) || err != v.err {
			t.Errorf("test [%v] failed on method ToUint32Array with params(initialArr: %v), expected to be %v got %v and error to be %v got %v",
				k, v.arr, v.response, *resArr, v.err, err)
		}
	}
}

func TestToUint64Array(t *testing.T) {
	type testData struct {
		arr      []int
		response []uint64
		err      error
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []int{}, response: []uint64{}, err: nil},
		"just one element":  testData{arr: []int{1}, response: []uint64{1}, err: nil},
		"multiple elements": testData{arr: []int{2, 3, 1}, response: []uint64{2, 3, 1}, err: nil},
	}

	for k, v := range examples {
		initialArr := IntArray(v.arr)

		resArr, err := initialArr.ToUint64Array()

		if resArr == nil || !reflect.DeepEqual(*resArr, v.response) || err != v.err {
			t.Errorf("test [%v] failed on method ToUint64Array with params(initialArr: %v), expected to be %v got %v and error to be %v got %v",
				k, v.arr, v.response, *resArr, v.err, err)
		}
	}
}

func TestToJson(t *testing.T) {
	arr := []int{2, 3, 4}
	response := "[2,3,4]"
	var errResp error

	initialArr := IntArray(arr)

	jsonStr, err := initialArr.ToJSON()

	if jsonStr != response || err != errResp {
		t.Errorf("test failed on method ToJSON with params(initialArr: %v), expected to be %v got %v and error to be %v got %v",
			arr, response, jsonStr, errResp, err)
	}

}
