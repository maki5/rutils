package arrays

type FloatArray []float64

// Delete deletes float element from array of floats
func (arr *FloatArray) Delete(elem float64) {
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

// Contains checks if array of floats contains provided float
func (arr *FloatArray) Contains(elem float64) bool {
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

// Clear remove all elements from array of floats
func (arr *FloatArray) Clear() {
	newArr := *arr
	*arr = newArr[:0]
}

// Collect return new array contained values returned by the provided function
func (arr *FloatArray) Collect(exec func(el float64) float64) []float64 {
	newArr := make([]float64, 0, 0)

	for _, e := range *arr {
		newArr = append(newArr, exec(e))
	}

	return newArr
}

// Concat append elements of other arrays to self
func (arr *FloatArray) Concat(arrays ...[]float64) {
	if len(arrays) == 0 {
		return
	}

	newArr := []float64(*arr)

	for _, a := range arrays {
		newArr = append(newArr, a...)
	}

	*arr = newArr
}

// Index return index of first matched float in array if not found return nil
func (arr *FloatArray) Index(elem float64) *int {
	if len(*arr) == 0 {
		return nil
	}

	newArr := []float64(*arr)

	for i, el := range newArr {
		if el == elem {
			return &i
		}
	}
	return nil
}

// Map return new array contained values returned by the provided function
func (arr *FloatArray) Map(exec func(el float64) float64) []float64 {
	return arr.Collect(exec)
}

// Min return min float64
func (arr *FloatArray) Min() *float64 {
	if len(*arr) == 0 {
		return nil
	}

	newArr := []float64(*arr)

	var min = newArr[0]

	for _, el := range newArr {
		if el < min {
			min = el
		}
	}

	return &min
}

// Max return max int
func (arr *FloatArray) Max() *float64 {
	if len(*arr) == 0 {
		return nil
	}

	newArr := []float64(*arr)

	var max = newArr[0]

	for _, el := range newArr {
		if el > max {
			max = el
		}
	}
	return &max
}

// Pop removes last element from array and returns it
func (arr *FloatArray) Pop(args ...int) *float64 {
	if len(*arr) == 0 {
		return nil
	}

	n := 1
	var last float64
	var newArr []float64

	if len(args) > 0 {
		n = args[0]
		if len(*arr) < n {
			return nil
		}
	}

	for i := 0; i < n; i++ {
		newArr = []float64(*arr)
		last = newArr[len(newArr)-1]
		arr.Delete(last)
	}

	return &last
}

// Push append element to array
func (arr *FloatArray) Push(elem float64) {
	newArr := []float64(*arr)
	newArr = append(newArr, elem)
	*arr = newArr
}

// Select returns a new array containing all elements of array for which the given block returns true
func (arr *FloatArray) Select(exec func(str float64) bool) []float64 {
	if len(*arr) == 0 {
		return []float64{}
	}

	newArr := []float64(*arr)
	resArr := make([]float64, 0, 0)

	for _, el := range newArr {
		if exec(el) {
			resArr = append(resArr, el)
		}
	}

	return resArr
}

// Uniq removes duplicated elements form given array
func (arr *FloatArray) Uniq() {
	if len(*arr) == 0 {
		return
	}

	newArr := []float64(*arr)
	strMap := make(map[float64]int)

	for _, el := range newArr {
		// value doesn't matter here cause we collect just keys
		strMap[el] = 1
	}

	resArr := make([]float64, 0, 0)
	for k := range strMap {
		resArr = append(resArr, k)
	}

	*arr = resArr
}
