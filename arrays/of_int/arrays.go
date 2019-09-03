package arrays

import (
	"encoding/json"
	"strconv"
)

type IntArray []int

// Delete deletes int element from array of ints
func (arr *IntArray) Delete(elem int) {
	// original version got form https://yourbasic.org/golang/delete-element-slice/
	var elemIndex *int

	for i, e := range *arr {
		if e == elem {
			elemIndex = &i
			break
		}
	}

	if elemIndex != nil {
		// Remove the element at index i from a.
		copy((*arr)[*elemIndex:], (*arr)[*elemIndex+1:]) // Shift a[i+1:] left one index.
		*arr = (*arr)[:len(*arr)-1]                      // Truncate slice.

	}
}

// Contains checks if array of ints contains provided int
func (arr *IntArray) Contains(elem int) bool {
	if len(*arr) == 0 {
		return false
	}

	for _, e := range *arr {
		if e == elem {
			return true
		}
	}

	return false
}

// Clear remove all elements from array of ints
func (arr *IntArray) Clear() {
	newArr := *arr
	*arr = newArr[:0]
}

// Collect return new array contained values returned by the provided function
func (arr *IntArray) Collect(exec func(el int) int) []int {
	newArr := make([]int, 0, 0)

	for _, e := range *arr {
		newArr = append(newArr, exec(e))
	}

	return newArr
}

// Concat append elements of other arrays to self
func (arr *IntArray) Concat(arrays ...[]int) {
	if len(arrays) == 0 {
		return
	}

	for _, a := range arrays {
		*arr = append(*arr, a...)
	}

}

// Index return index of first matched int in array if not found return nil
func (arr *IntArray) Index(elem int) *int {
	if len(*arr) == 0 {
		return nil
	}

	for i, el := range *arr {
		if el == elem {
			return &i
		}
	}
	return nil
}

// Map return new array contained values returned by the provided function
func (arr *IntArray) Map(exec func(el int) int) []int {
	return arr.Collect(exec)
}

// Min return min int
func (arr *IntArray) Min() *int {
	if len(*arr) == 0 {
		return nil
	}

	newArr := []int(*arr)

	var min = newArr[0]

	for _, el := range newArr {
		if el < min {
			min = el
		}
	}

	return &min
}

// Max return max int
func (arr *IntArray) Max() *int {
	if len(*arr) == 0 {
		return nil
	}

	newArr := []int(*arr)

	var max = newArr[0]

	for _, el := range newArr {
		if el > max {
			max = el
		}
	}
	return &max
}

// Pop removes last element from array and returns it
func (arr *IntArray) Pop(args ...int) *int {
	if len(*arr) == 0 {
		return nil
	}

	n := 1
	var last int
	var newArr []int

	if len(args) > 0 {
		n = args[0]
		if len(*arr) < n {
			return nil
		}
	}

	for i := 0; i < n; i++ {
		newArr = []int(*arr)
		last = newArr[len(newArr)-1]
		arr.Delete(last)
	}

	return &last
}

// Push append element to array
func (arr *IntArray) Push(elem int) {
	*arr = append(*arr, elem)
}

// Select returns a new array containing all elements of array for which the given block returns true
func (arr *IntArray) Select(exec func(str int) bool) []int {
	if len(*arr) == 0 {
		return []int{}
	}

	resArr := make([]int, 0, 0)

	for _, el := range *arr {
		if exec(el) {
			resArr = append(resArr, el)
		}
	}

	return resArr
}

// Uniq removes duplicated elements form given array
func (arr *IntArray) Uniq() {
	if len(*arr) == 0 {
		return
	}

	strMap := make(map[int]int)

	for _, el := range *arr {
		// value doesn't matter here cause we collect just keys
		strMap[el] = 1
	}

	resArr := make([]int, 0, 0)
	for k := range strMap {
		resArr = append(resArr, k)
	}

	*arr = resArr
}

// ToStringArray implements Convertible for converting to string array
func (arr *IntArray) ToStringArray() (*[]string, error) {
	newArr := make([]string, 0, 0)

	for _, el := range *arr {
		newArr = append(newArr, strconv.Itoa(el))
	}
	return &newArr, nil
}

// ToFloat64Array implements Convertible for converting to float32 array
func (arr *IntArray) ToFloat64Array() (*[]float64, error) {
	newArr := make([]float64, 0, 0)

	for _, el := range *arr {
		newArr = append(newArr, float64(el))
	}
	return &newArr, nil
}

// ToFloat32Array implements Convertible for converting to float32 array
func (arr *IntArray) ToFloat32Array() (*[]float32, error) {
	newArr := make([]float32, 0, 0)

	for _, el := range *arr {
		newArr = append(newArr, float32(el))
	}
	return &newArr, nil
}

// ToInt64Array implements Convertible for converting to int64 array
func (arr *IntArray) ToInt64Array() (*[]int64, error) {
	newArr := make([]int64, 0, 0)

	for _, el := range *arr {
		newArr = append(newArr, int64(el))
	}
	return &newArr, nil
}

// ToInt32Array implements Convertible for converting to int32 array
func (arr *IntArray) ToInt32Array() (*[]int32, error) {
	newArr := make([]int32, 0, 0)

	for _, el := range *arr {
		newArr = append(newArr, int32(el))
	}
	return &newArr, nil
}

// ToUintArray implements Convertible for converting to uint array
func (arr *IntArray) ToUintArray() (*[]uint, error) {
	newArr := make([]uint, 0, 0)

	for _, el := range *arr {
		newArr = append(newArr, uint(el))
	}
	return &newArr, nil
}

// ToUint32Array implements Convertible for converting to uint32 array
func (arr *IntArray) ToUint32Array() (*[]uint32, error) {
	newArr := make([]uint32, 0, 0)

	for _, el := range *arr {
		newArr = append(newArr, uint32(el))
	}
	return &newArr, nil
}

// ToUint64Array implements Convertible for converting to uint64 array
func (arr *IntArray) ToUint64Array() (*[]uint64, error) {
	newArr := make([]uint64, 0, 0)

	for _, el := range *arr {
		newArr = append(newArr, uint64(el))
	}
	return &newArr, nil
}

// ToJSON implements Convertible for converting to json string
func (arr *IntArray) ToJSON() (string, error) {
	data, err := json.Marshal(*arr)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
