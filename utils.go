package rutils

func IntPtr(i int) *int {
	return &i
}

func InverseInt(n int) int{
	return n - (n*2)
}