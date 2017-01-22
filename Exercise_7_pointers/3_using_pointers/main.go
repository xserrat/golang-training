package main

import "fmt"

func main() {
	a_number := 34

	fmt.Println("Value: ", a_number, " Memory address: ", &a_number)

	var a_pointer *int = &a_number

	fmt.Println("Pointer memory address: ", a_pointer, " Value stored at the memory address: ", *a_pointer)

	*a_pointer = 25

	fmt.Println("Now the pointer with memory address: ", a_pointer, " stores the value: ", *a_pointer)
	fmt.Println("And due to the memory address is the same as 'a_number' variable, now this variable contains the value: ", a_number)
}
