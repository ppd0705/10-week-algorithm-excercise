package main

func toLowerCase(str string) string {
	bytes := []byte(str)
	for i := 0; i < len(bytes); i++ {
		if 'A' <= bytes[i] && bytes[i] <= 'Z' {
			bytes[i] += 32
		}
	}
	return string(bytes)
}
