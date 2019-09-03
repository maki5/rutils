package arrays

import (
	"mlab/rutils"
	"reflect"
	"sort"
	"testing"
)

func TestDelete(t *testing.T) {
	type testData struct {
		arr      []float64
		selector float64
		response []float64
	}
	examples := map[string]testData{
		"empty arr":                    testData{arr: []float64{}, selector: 1, response: []float64{}},
		"sample int arr":               testData{arr: []float64{1, 2, 3, 4}, selector: 1, response: []float64{2, 3, 4}},
		"sample int arr 1":             testData{arr: []float64{1, 2, 3, 4}, selector: 2, response: []float64{1, 3, 4}},
		"int doesn't present in array": testData{arr: []float64{1, 2, 3, 4}, selector: 5, response: []float64{1, 2, 3, 4}},
	}

	badExamples := map[string]testData{
		"param doesn't exist": testData{arr: []float64{1, 2, 3, 4}, selector: 5, response: []float64{1, 2, 3, 4, 5}},
		"param not deleted":   testData{arr: []float64{1, 2, 3, 4}, selector: 1, response: []float64{1, 2, 3, 4}},
	}

	for k, v := range examples {
		initialArr := FloatArray(v.arr)

		initialArr.Delete(v.selector)
		if !reflect.DeepEqual([]float64(initialArr), v.response) {
			t.Errorf("test [%v] failed on method Delete with params(initialArr: %v, selector %v), expected %v got %v",
				k, v.arr, v.selector, v.response, initialArr)
		}
	}

	for k, v := range badExamples {
		initialArr := FloatArray(v.arr)

		initialArr.Delete(v.selector)
		if reflect.DeepEqual([]float64(initialArr), v.response) {
			t.Errorf("test [%v] failed on method Delete with params(initialArr: %v, selector %v), expected to not be %v",
				k, v.arr, v.selector, v.response)
		}
	}
}

func TestContains(t *testing.T) {
	type testData struct {
		arr      []float64
		selector float64
		response bool
	}

	examples := map[string]testData{
		"empty arr":                    testData{arr: []float64{}, selector: 1, response: false},
		"sample str arr":               testData{arr: []float64{1, 2, 3, 4}, selector: 1, response: true},
		"sample str arr 1":             testData{arr: []float64{1, 2, 3, 4}, selector: 2, response: true},
		"str doesn't present in array": testData{arr: []float64{1, 2, 3, 4}, selector: 5, response: false},
	}

	for k, v := range examples {
		arr := FloatArray(v.arr)
		resp := arr.Contains(v.selector)
		if v.response != resp {
			t.Errorf("test [%v] failed on method Contains with params(initialArr: %v, selector %v), expected %v got %v",
				k, v.arr, v.selector, v.response, resp)
		}
	}
}

func TestClear(t *testing.T) {
	type testData struct {
		arr []float64
	}

	examples := map[string]testData{
		"empty arr":        testData{arr: []float64{}},
		"sample str arr":   testData{arr: []float64{1, 2, 3, 4}},
		"sample str arr 1": testData{arr: []float64{1}},
	}

	for k, v := range examples {
		initialArr := FloatArray(v.arr)

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
		arr      []float64
		execFunc func(el float64) float64
		response []float64
	}

	f1 := func(el float64) float64 {
		return el + 1
	}

	f2 := func(el float64) float64 {
		if el != 0 {
			return el + 1
		}

		return -1
	}

	examples := map[string]testData{
		"empty arr":        testData{arr: []float64{}, execFunc: f1, response: []float64{}},
		"sample str arr":   testData{arr: []float64{1, 2, 3, 4}, execFunc: f1, response: []float64{2, 3, 4, 5}},
		"sample str arr 1": testData{arr: []float64{1, 0, 2, 0, 3}, execFunc: f2, response: []float64{2, -1, 3, -1, 4}},
	}

	for k, v := range examples {
		arr := FloatArray(v.arr)
		newArr := arr.Collect(v.execFunc)

		if !reflect.DeepEqual(newArr, v.response) {
			t.Errorf("test [%v] failed on method Collect with params(initialArr: %v), expected to be %v got %v",
				k, v.arr, v.response, newArr)
		}
	}
}

func TestConcat(t *testing.T) {
	type testData struct {
		arr            []float64
		arraysToConcat [][]float64
		response       []float64
	}

	examples := map[string]testData{
		"empty arr": testData{arr: []float64{}, arraysToConcat: [][]float64{}, response: []float64{}},
		"concat one array": testData{arr: []float64{1, 4}, arraysToConcat: [][]float64{[]float64{2, 3}},
			response: []float64{1, 4, 2, 3}},
		"concat multiple arrays": testData{arr: []float64{1, 4}, arraysToConcat: [][]float64{[]float64{2, 3}, []float64{5}},
			response: []float64{1, 4, 2, 3, 5}},
	}

	for k, v := range examples {
		initialArr := FloatArray(v.arr)
		initialArr.Concat(v.arraysToConcat...)

		if !reflect.DeepEqual([]float64(initialArr), v.response) {
			t.Errorf("test [%v] failed on method Concat with params(initialArr: %v), expected to be %v got %v",
				k, v.arr, v.response, initialArr)
		}
	}
}

func TestIndex(t *testing.T) {
	type testData struct {
		arr      []float64
		selector float64
		response *int
	}

	examples := map[string]testData{
		"empty arr":                  testData{arr: []float64{}, selector: 1, response: nil},
		"one occurence in arr":       testData{arr: []float64{1, 4}, selector: 1, response: rutils.IntPtr(0)},
		"multiple occurences in arr": testData{arr: []float64{1, 4, 1}, selector: 1, response: rutils.IntPtr(0)},
	}

	for k, v := range examples {
		initialArr := FloatArray(v.arr)
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
		arr      []float64
		response *float64
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []float64{}, response: nil},
		"one element":       testData{arr: []float64{1}, response: rutils.FloatPtr(1)},
		"multiple elements": testData{arr: []float64{1, 3, 4, 7, 3, 2}, response: rutils.FloatPtr(1)},
	}

	for k, v := range examples {
		initialArr := FloatArray(v.arr)
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
		arr      []float64
		response *float64
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []float64{}, response: nil},
		"one element":       testData{arr: []float64{1}, response: rutils.FloatPtr(1)},
		"multiple elements": testData{arr: []float64{1, 3, 4, 7, 3, 2}, response: rutils.FloatPtr(7)},
	}

	for k, v := range examples {
		initialArr := FloatArray(v.arr)
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
		arr      []float64
		newArr   []float64
		n        *int
		response *float64
	}

	examples := map[string]testData{
		"empty arr":         testData{arr: []float64{}, newArr: []float64{}, response: nil},
		"one element":       testData{arr: []float64{1}, newArr: []float64{}, response: rutils.FloatPtr(1)},
		"multiple elements": testData{arr: []float64{1, 2, 3}, newArr: []float64{1, 2}, response: rutils.FloatPtr(3)},
		"pop two elements":  testData{arr: []float64{1, 2, 3}, newArr: []float64{1}, response: rutils.FloatPtr(2), n: rutils.IntPtr(2)},
		"pop zero elements": testData{arr: []float64{1, 2, 3}, newArr: []float64{1, 2, 3}, response: nil, n: rutils.IntPtr(0)},
	}

	for k, v := range examples {
		initialArr := FloatArray(v.arr)
		var resp *float64
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

		if !reflect.DeepEqual([]float64(initialArr), v.newArr) {
			t.Errorf("test [%v] failed on method Pop with params(initialArr: %v), expected to be %v got %v",
				k, v.arr, initialArr, v.newArr)
		}
	}

}

func TestSelect(t *testing.T) {
	type testData struct {
		arr      []float64
		execFunc func(elem float64) bool
		response []float64
	}

	f1 := func(elem float64) bool {
		if elem == 1 {
			return true
		}
		return false
	}

	f2 := func(elem float64) bool {
		if elem == 3 || elem == 2 {
			return true
		}

		return false
	}

	examples := map[string]testData{
		"empty arr":          testData{arr: []float64{}, execFunc: f1, response: []float64{}},
		"func returns false": testData{arr: []float64{2}, execFunc: f1, response: []float64{}},
		"valid elements":     testData{arr: []float64{2, 3, 1}, execFunc: f2, response: []float64{2, 3}},
	}

	for k, v := range examples {
		initialArr := FloatArray(v.arr)

		resp := initialArr.Select(v.execFunc)

		if !reflect.DeepEqual(resp, v.response) {
			t.Errorf("test [%v] failed on method Select with params(initialArr: %v), expected to be %v got %v",
				k, v.arr, v.response, resp)
		}

	}
}

func TestUniq(t *testing.T) {
	type testData struct {
		arr      []float64
		response []float64
	}

	examples := map[string]testData{
		"empty arr":                  testData{arr: []float64{}, response: []float64{}},
		"just one element":           testData{arr: []float64{1}, response: []float64{1}},
		"multiple uniq elements":     testData{arr: []float64{2, 3, 1}, response: []float64{2, 3, 1}},
		"multiple non uniq elements": testData{arr: []float64{2, 3, 1, 2, 1}, response: []float64{2, 3, 1}},
	}

	for k, v := range examples {
		initialArr := FloatArray(v.arr)

		initialArr.Uniq()

		sort.Float64s([]float64(initialArr))
		sort.Float64s(v.response)

		if !reflect.DeepEqual([]float64(initialArr), v.response) {
			t.Errorf("test [%v] failed on method Uniq with params(initialArr: %v), expected to be %v got %v",
				k, v.arr, v.response, initialArr)
		}

	}

}
