package arrays

import (
	arrayOfFloats "mlab/rutils/arrays/of_float64"
	arrayOfInts "mlab/rutils/arrays/of_int"
	arrayOfStrings "mlab/rutils/arrays/of_string"
	"testing"
)

func TestConvertible(t *testing.T) {
	sa := arrayOfStrings.StringArray{}
	ia := arrayOfInts.IntArray{}
	fa := arrayOfFloats.FloatArray{}

	acceptAndCheckConvertible(&sa)
	acceptAndCheckConvertible(&ia)
	acceptAndCheckConvertible(&fa)
}

func acceptAndCheckConvertible(c Convertible) bool {
	return true
}
