package main

import "testing"

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
