package main

func IsNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func IsLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}
