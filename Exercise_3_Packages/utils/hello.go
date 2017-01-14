package utils

import "fmt"

func Hello(user string) {
	sayHello(user)
}

func sayHello(user string) {
	fmt.Println("Hello ", user, "!!")
}
