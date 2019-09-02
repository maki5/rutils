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
