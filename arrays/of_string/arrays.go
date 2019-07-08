package arrays

// DeleteString : deletes string element from array of strings
func DeleteString(arr *[]string, elem string) {
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

func Contains(arr *[]string, elem string) bool {
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

func Clear(arr *[]string) {
	newArr := *arr
	newArr = newArr[:0]
	arr = &newArr
}
