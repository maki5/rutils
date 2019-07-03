package arrays

// DeleteString : original version got form https://yourbasic.org/golang/delete-element-slice/
func DeleteString(arr *[]string, elem string) {
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
		(*arr)[len(*arr)-1] = ""                       // Erase last element (write zero value).
		*arr = (*arr)[:len(*arr)-1]                 // Truncate slice.

	}
}
