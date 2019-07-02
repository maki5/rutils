package strings

import (
	"fmt"
)

func At(str string, pos interface{}) (string, error) {
	if p, ok := pos.(int); ok {
		return string(str[p]), nil
	}

	if p, ok := pos.([]int); ok {
		if len(p) != 2 {
			return "", fmt.Errorf("wrong params, expected 2 got %v", len(p))
		}

		return string(str[p[0] : p[1]+1]), nil
	}
	return "", fmt.Errorf("wrong params")
}
