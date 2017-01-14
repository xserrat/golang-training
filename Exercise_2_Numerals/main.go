package main

import "fmt"

func main() {
	for i := 65; i <= 80; i++ {
		fmt.Printf("- Decimal: %d \t Binary: %b \t Hexadecimal: %x \t UTF-8: %q \n", i, i, i, i)
	}
}
