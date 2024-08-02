package fib

import "testing"

func BenchmarkFib10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibRecursive(25)
	}
}
