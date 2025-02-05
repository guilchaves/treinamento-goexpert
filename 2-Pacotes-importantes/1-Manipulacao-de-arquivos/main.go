//INFO: 2.1-Manipulacao de arquivos
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Write file
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	//size, err := f.WriteString("Hello, World!\n")
	size, err := f.Write([]byte("Hello, World!\n"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Wrote %d bytes\n", size)
	f.Close()

	// Read file
	file, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))

	// Read file with bufio
	file2, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file2)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Print(string(buffer[:n]))
	}
}
