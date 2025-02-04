//INFO: 10-Closures
package main

import (
	"fmt"
)

// Closure function
func main() {
    total := func() int{
        return sum(1, 430, 23, 55) * 2
    }()

    fmt.Println(total)

}

func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
