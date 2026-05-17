package main

func IsNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func IsLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// because go dosen't have ternary operator
func If[T any](condition bool, a, b T) T {
	if condition {
		return a
	}
	return b
}
