package arrays

import (
	"encoding/json"
	"strconv"

	"github.com/maki5/rutils"
)

// StringArray alias type for []string
type StringArray []string

// Delete deletes string element from array of strings
func (arr *StringArray) Delete(elem string) {
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
		(*arr)[len(*arr)-1] = ""                         // Erase last element (write zero value).
		*arr = (*arr)[:len(*arr)-1]                      // Truncate slice.

	}
}

// Contains checks if array of strings contains provided string
func (arr *StringArray) Contains(elem string) bool {
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

// Clear remove all elements from array of strings
func (arr *StringArray) Clear() {
	newArr := *arr
	*arr = newArr[:0]
}

// Collect return new array contained values returned by the provided function
func (arr *StringArray) Collect(exec func(el string) string) []string {
	newArr := make([]string, 0, 0)

	for _, e := range *arr {
		newArr = append(newArr, exec(e))
	}

	return newArr
}

// Compact removes all empty elements from given array
func (arr *StringArray) Compact() {
	newArr := make([]string, 0, 0)

	for _, e := range *arr {
		if !rutils.Blank(e) {
			newArr = append(newArr, e)
		}
	}

	*arr = newArr
}

// Concat append elements of other arrays to self
func (arr *StringArray) Concat(arrays ...[]string) {
	if len(arrays) == 0 {
		return
	}

	for _, a := range arrays {
		*arr = append(*arr, a...)
	}

}

// Index return index of first matched string in array if not found return -1
func (arr *StringArray) Index(elem string) int {
	if len(*arr) == 0 || rutils.Blank(elem) {
		return -1
	}

	for i, el := range *arr {
		if el == elem {
			return i
		}
	}
	return -1
}

// Map return new array contained values returned by the provided function
func (arr *StringArray) Map(exec func(el string) string) []string {
	return arr.Collect(exec)
}

// Min return min string
func (arr *StringArray) Min() string {
	if len(*arr) == 0 {
		return ""
	}

	newArr := []string(*arr)

	var min = newArr[0]

	for _, el := range newArr {
		if len(el) < len(min) {
			min = el
		} else if len(el) == len(min) {
			if stringWeight(el) < stringWeight(min) {
				min = el
			}
		}
	}

	return min
}

// Max return max string
func (arr *StringArray) Max() string {
	if len(*arr) == 0 {
		return ""
	}

	newArr := []string(*arr)

	var max = newArr[0]

	for _, el := range newArr {
		if len(el) > len(max) {
			max = el
		} else if len(el) == len(max) {
			if stringWeight(el) > stringWeight(max) {
				max = el
			}
		}
	}
	return max
}

// Pop removes last element from array and returns it
func (arr *StringArray) Pop(args ...int) string {
	if len(*arr) == 0 {
		return ""
	}

	n := 1
	var last string
	var newArr []string

	if len(args) > 0 {
		n = args[0]
		if len(*arr) < n {
			return ""
		}
	}

	for i := 0; i < n; i++ {
		newArr = []string(*arr)
		last = newArr[len(newArr)-1]
		arr.Delete(last)
	}

	return last
}

// Push append element to array
func (arr *StringArray) Push(elem string) {
	*arr = append(*arr, elem)

}

// Select returns a new array containing all elements of array for which the given block returns true
func (arr *StringArray) Select(exec func(elem string) bool) []string {
	if len(*arr) == 0 {
		return []string{}
	}

	resArr := make([]string, 0, 0)

	for _, el := range *arr {
		if exec(el) {
			resArr = append(resArr, el)
		}
	}

	return resArr
}

// Uniq removes duplicated elements form given array
func (arr *StringArray) Uniq() {
	if len(*arr) == 0 {
		return
	}

	strMap := make(map[string]int)

	for _, el := range *arr {
		// value doesn't matter here cause we collect just keys
		strMap[el] = 1
	}

	resArr := make([]string, 0, 0)
	for k := range strMap {
		resArr = append(resArr, k)
	}

	*arr = resArr
}

// ToStringArray implements Convertible for converting to string array
func (arr *StringArray) ToStringArray() (*[]string, error) {
	newArr := []string(*arr)
	return &newArr, nil
}

// ToFloat64Array implements Convertible for converting to float32 array
func (arr *StringArray) ToFloat64Array() (*[]float64, error) {
	newArr := make([]float64, 0, 0)

	for _, el := range *arr {
		f, err := strconv.ParseFloat(el, 64)
		if err != nil {
			return nil, err
		}
		newArr = append(newArr, f)
	}
	return &newArr, nil
}

// ToFloat32Array implements Convertible for converting to float32 array
func (arr *StringArray) ToFloat32Array() (*[]float32, error) {
	newArr := make([]float32, 0, 0)

	for _, el := range *arr {
		f, err := strconv.ParseFloat(el, 64)
		if err != nil {
			return nil, err
		}
		newArr = append(newArr, float32(f))
	}
	return &newArr, nil
}

// ToInt64Array implements Convertible for converting to int64 array
func (arr *StringArray) ToInt64Array() (*[]int64, error) {
	newArr := make([]int64, 0, 0)

	for _, el := range *arr {
		i, err := strconv.ParseInt(el, 10, 64)
		if err != nil {
			return nil, err
		}

		newArr = append(newArr, i)
	}
	return &newArr, nil
}

// ToInt32Array implements Convertible for converting to int32 array
func (arr *StringArray) ToInt32Array() (*[]int32, error) {
	newArr := make([]int32, 0, 0)

	for _, el := range *arr {
		i, err := strconv.ParseInt(el, 10, 64)
		if err != nil {
			return nil, err
		}

		newArr = append(newArr, int32(i))
	}
	return &newArr, nil
}

// ToUintArray implements Convertible for converting to uint array
func (arr *StringArray) ToUintArray() (*[]uint, error) {
	newArr := make([]uint, 0, 0)

	for _, el := range *arr {
		i, err := strconv.ParseUint(el, 10, 64)
		if err != nil {
			return nil, err
		}
		newArr = append(newArr, uint(i))
	}
	return &newArr, nil
}

// ToUint32Array implements Convertible for converting to uint32 array
func (arr *StringArray) ToUint32Array() (*[]uint32, error) {
	newArr := make([]uint32, 0, 0)

	for _, el := range *arr {
		i, err := strconv.ParseUint(el, 10, 64)
		if err != nil {
			return nil, err
		}
		newArr = append(newArr, uint32(i))
	}
	return &newArr, nil
}

// ToUint64Array implements Convertible for converting to uint64 array
func (arr *StringArray) ToUint64Array() (*[]uint64, error) {
	newArr := make([]uint64, 0, 0)

	for _, el := range *arr {
		i, err := strconv.ParseUint(el, 10, 64)
		if err != nil {
			return nil, err
		}
		newArr = append(newArr, i)
	}
	return &newArr, nil
}

// ToJSON implements Convertible for converting to json string
func (arr *StringArray) ToJSON() (string, error) {
	data, err := json.Marshal(*arr)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// internal functions
func stringWeight(str string) int {
	sum := 0

	for _, el := range str {
		sum += int(el)
	}

	return sum
}
