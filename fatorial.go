package main

// Fatorial retorna o fatorial de um número
func Fatorial(x int) int {
	if x == 1 {
		return x
	}
	return x * Fatorial(x-1)
}
