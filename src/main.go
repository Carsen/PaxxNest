package main

import (
	"fmt"
	"os"

	"github.com/Carsen/PaxxNest/Login"
	"github.com/Carsen/PaxxNest/Manager"
)

func main() {
	switch Login.Login(true) {
	case true:
		fmt.Printf("Hello!")
		mgr := Manager.NewManager()
		mgr.AddManager("snap", Manager.SnapManager{})
		mgr.ListPackages()
	case false:
		fmt.Println("Goodbye!")
		os.Exit(1)
	}
}
