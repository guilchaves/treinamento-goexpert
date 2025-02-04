//INFO: 15-Quando usar ponteiros
package main

func sum(a, b *int) int {
    *a = 50
	return *a + *b
}

func main() {
    a := 10
    b := 20

    sum(&a, &b)
    println(a)
}
