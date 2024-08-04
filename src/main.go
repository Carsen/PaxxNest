package main

import (
	"fmt"
	"os"

	"github.com/Carsen/PaxxNest/Login"
	"github.com/Carsen/PaxxNest/Manager"
)

func main() {
	for Login.Login(true) {
		fmt.Println("Hello!\n")
		mgr := Manager.NewManager()
		mgr.AddManager("brew", Manager.BrewMan{})
		mgr.AddManager("snap", Manager.SnapMan{})

		fmt.Print("Press 'L' for list, 'I' for install, and 'R' for remove")
		var choice string
		fmt.Scanln(&choice)

		if choice == "L" || choice == "l" {
			Manager.ListInstalledPackages()
		}
		if choice == "i" || choice == "I" {
			var whichpkg string
			fmt.Scanln(&whichpkg)
			Manager.InstallPackage(whichpkg)
		}
	}
	for !Login.Login(true) {
		fmt.Println("Goodbye!")
		os.Exit(1)
	}
}
