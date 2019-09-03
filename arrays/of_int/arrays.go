package arrays

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

	newArr := []int(*arr)

	for _, a := range arrays {
		newArr = append(newArr, a...)
	}

	*arr = newArr
}

// Index return index of first matched int in array if not found return nil
func (arr *IntArray) Index(elem int) *int {
	if len(*arr) == 0 {
		return nil
	}

	newArr := []int(*arr)

	for i, el := range newArr {
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
	newArr := []int(*arr)
	newArr = append(newArr, elem)
	*arr = newArr
}

// Select returns a new array containing all elements of array for which the given block returns true
func (arr *IntArray) Select(exec func(str int) bool) []int {
	if len(*arr) == 0 {
		return []int{}
	}

	newArr := []int(*arr)
	resArr := make([]int, 0, 0)

	for _, el := range newArr {
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

	newArr := []int(*arr)
	strMap := make(map[int]int)

	for _, el := range newArr {
		// value doesn't matter here cause we collect just keys
		strMap[el] = 1
	}

	resArr := make([]int, 0, 0)
	for k := range strMap {
		resArr = append(resArr, k)
	}

	*arr = resArr
}
