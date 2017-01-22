package main

import "fmt"

func main() {
	a_variable := 22

	var a_pointer_variable *int = &a_variable

	fmt.Println(a_variable)
	fmt.Println(&a_variable)
	fmt.Println(a_pointer_variable)
}
