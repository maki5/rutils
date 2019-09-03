package arrays

import (
	"testing"

	arrayOfFloats "github.com/maki5/rutils/arrays/of_float64"
	arrayOfInts "github.com/maki5/rutils/arrays/of_int"
	arrayOfStrings "github.com/maki5/rutils/arrays/of_string"
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
