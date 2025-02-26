// INFO: 22-For
package main

func main() {

	for i := 0; i < 10; i++ {
		println(i)
	}

	nums := []string{"one", "two", "three", "four", "five"}
	for k, v := range nums {
		println(k, v)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}

}
