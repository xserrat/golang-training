package main

import (
	"fmt"
	"github.com/xserrat/golang-training/Exercise_6_memory_address/memory"
)

func main() {
	fmt.Println("Showing memory address:")

	a_variable := 13
	memory.ShowMemoryAddress(a_variable)

	fmt.Print("Introduce an input value to be printed then: ")

	var variable_with_input_data_stored string
	data_introduced := memory.PutTheInputDataFromKeyboardToTheVariableAndGetIt(variable_with_input_data_stored)

	fmt.Printf("The value introduced is '%s'\n", data_introduced)
}
