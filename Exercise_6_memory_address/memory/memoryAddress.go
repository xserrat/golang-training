package memory

import "fmt"

func ShowMemoryAddress(some_variable int) {
	fmt.Println("A value of the variable: ", some_variable)
	fmt.Printf("The memory address is: %d (%#x) \n", &some_variable, &some_variable)
}

func PutTheInputDataFromKeyboardToTheVariableAndGetIt(variable_to_store_input_data string) string {
	fmt.Scan(&variable_to_store_input_data)
	return variable_to_store_input_data
}