package main

import (
	"fmt"

	"github.com/Carsen/PaxxNest/Login"
)

func main() {
	switch Login.Login(true) {
	case true:
		fmt.Printf("Hello!")
	case false:
		fmt.Println("Goodbye!")
	}
}
