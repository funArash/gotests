package main

import (
	"testing"
)

func incnIntStatic(i *int, n int) {
	for k := 0; k < n; k++ {
		*i = *i + 1
	}
}

func BenchmarkStaticTypeOnHeap(b *testing.B) {
	i := new(int)
	incnIntStatic(i, b.N)
}

func BenchmarkStaticTypeOnStack(b *testing.B) {
	i := 0
	incnIntStatic(&i, b.N)
}
