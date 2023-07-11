package mymath

func Sum(a int, b int) (result int) {
	result = a + b
	return result
}

func Substract(a int, b int) (result int) {
	result = a - b
	return result
}

func Multiply(a int, b int) (result int) {
	result = a * b
	return result
}

func Divide(a int, b int) (result float32) {
	result = float32(a) / float32(b)
	return result
}
