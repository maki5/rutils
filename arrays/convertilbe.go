package arrays

// Convertible interface is used for converting arrays of different types
type Convertible interface {
	ToStringArray() (*[]string, error)
	ToFloat32Array() (*[]float32, error)
	ToFloat64Array() (*[]float64, error)
	ToInt64Array() (*[]int64, error)
	ToInt32Array() (*[]int32, error)
	ToUintArray() (*[]uint, error)
	ToUint32Array() (*[]uint32, error)
	ToUint64Array() (*[]uint64, error)
	ToJSON() (string, error)
}
