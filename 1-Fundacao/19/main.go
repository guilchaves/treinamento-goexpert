// INFO: 19-Generics
package main

type MyNumber int

type Number interface {
	~int | float64
}

func Sum[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

func isEqual[T comparable](a, b T) bool {
	return a == b
}

func main() {
	m := map[string]int{"Guilherme": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Guilherme": 1000.20, "João": 2000.3, "Maria": 3000.1}
	m3 := map[string]MyNumber{"Guilherme": 1000, "João": 2000, "Maria": 3000}

	println(Sum(m))
	println(Sum(m2))
	println(Sum(m3))
}
