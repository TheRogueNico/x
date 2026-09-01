package main

import "fmt"

func rotN(s string, n int) string {
	// Normalize N
	n = ((n % 26) + 26) % 26

	b := []byte(s)
	out := make([]byte, len(b))

	for i, c := range b {
		switch {
		case c >= 'a' && c <= 'z':
			out[i] = 'a' + (c-'a'+byte(n))%26
		case c >= 'A' && c <= 'Z':
			out[i] = 'A' + (c-'A'+byte(n))%26
		default:
			out[i] = c
		}
	}

	return string(out)
}

func main() {
	fmt.Println(rotN("Uryyb, Jbeyq!", 13))
}
