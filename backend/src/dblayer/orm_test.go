package dblayer

import "testing"

func BenchmarkHashPassword(b *testing.B) {
	text := "Hello world"
	for x := 0; x < b.N; x++ {
		hashPassword(&text)
	}
}
