package main

// Fatorial retorna o fatorial de um número
func Fatorial(x int) int {
	fatorial := x
	if x < 0 {
		return 0
	}
	if x > 0 {
		for f := x - 1; f > 0; f-- {
			fatorial *= f
		}
	}
	return fatorial
}
