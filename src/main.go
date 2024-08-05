package main

import (
	"fmt"

	"github.com/Carsen/PaxxNest/Login"
	"github.com/Carsen/PaxxNest/Manager"
)

func main() {
	switch Login.Login(true) {
	case true:
		fmt.Println("Hello!\n")
		mgr := Manager.NewManager()
		//	mgr.AddManager("brew", Manager.BrewMan{}) // Uncomment this line for Homebrew integration
		//	mgr.AddManager("snap", Manager.SnapMan{}) // Uncomment this line for Snap integration
		//	mgr.AddManager("pacman", Manager.PacMan{}) // Uncomment this line for Pacman integration
		//	mgr.AddManager("winget", Manager.WingetMan{}) // Uncomment this line for Winget integration

		fmt.Println("Press 'L' for list, 'I' for install, and 'R' for remove")
		var choice string
		fmt.Scanln(">", &choice)

		if choice == "L" || choice == "l" {
			mgr.ListPackages()
		}
		if choice == "i" || choice == "I" {
			var whichpkg string
			fmt.Scanln(&whichpkg)
			mgr.InstallPackage(whichpkg)
		}
		if choice == "r" || choice == "R" {
			var whichpkg string
			fmt.Scanln(&whichpkg)
			mgr.RemovePackage(whichpkg)
		}
	}
}
