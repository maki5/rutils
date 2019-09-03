package arrays

import (
	"mlab/rutils"
)

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

	newArr := []string(*arr)

	for _, a := range arrays {
		newArr = append(newArr, a...)
	}

	*arr = newArr
}

// Index return index of first matched string in array if not found return -1
func (arr *StringArray) Index(elem string) int {
	if len(*arr) == 0 || rutils.Blank(elem) {
		return -1
	}

	newArr := []string(*arr)

	for i, el := range newArr {
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
	newArr := []string(*arr)
	newArr = append(newArr, elem)
	*arr = newArr
}

// Select returns a new array containing all elements of array for which the given block returns true
func (arr *StringArray) Select(exec func(elem string) bool) []string {
	if len(*arr) == 0 {
		return []string{}
	}

	newArr := []string(*arr)
	resArr := make([]string, 0, 0)

	for _, el := range newArr {
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

	newArr := []string(*arr)
	strMap := make(map[string]int)

	for _, el := range newArr {
		// value doesn't matter here cause we collect just keys
		strMap[el] = 1
	}

	resArr := make([]string, 0, 0)
	for k := range strMap {
		resArr = append(resArr, k)
	}

	*arr = resArr
}

// internal functions
func stringWeight(str string) int {
	sum := 0

	for _, el := range str {
		sum += int(el)
	}

	return sum
}
