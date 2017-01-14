package main

import "fmt"

const Value int = 22

func main() {
	const Another_value int = 44
	fmt.Println(Value, Another_value)
}