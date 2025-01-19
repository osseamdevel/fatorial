package main

import (
	"fmt"
	"testing"
)

func TestFatorial(t *testing.T) {
	resultado := Fatorial(5)
	esperado := 120

	if resultado != esperado {
		t.Errorf("resultado %d, esperado %d", resultado, esperado)
	}
}

func BenchmarkFatorial(b *testing.B) {
	for f := 0; f < b.N; f++ {
		Fatorial(7)
	}
}

func ExampleFatorial() {
	fatorial := Fatorial(10)
	fmt.Println(fatorial)
	// Output: 3628800
}
