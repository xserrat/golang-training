package main

import "fmt"

func main() {
	value := 33

	var pointer *int = &value

	fmt.Println("This memory pointer has the value: ", *pointer, " stored in the memory address: ", pointer)
}
