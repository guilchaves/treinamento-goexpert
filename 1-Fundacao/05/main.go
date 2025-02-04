// INFO: 5-Percorrendo arrays
package main

import "fmt"

func main() {
	var arr [3]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for i, v := range arr {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

}
