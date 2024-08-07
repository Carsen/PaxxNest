package main

import (
	"fmt"
	
	"github.com/Carsen/PaxxNest/Manager"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	mgr := Manager.NewManager()
	mgr.AddManager("brew", Manager.BrewMan{})        //	Uncomment this line for Homebrew integration
	mgr.AddManager("snap", Manager.SnapMan{})        //	Uncomment this line for Snap integration
	mgr.AddManager("pacman", Manager.PacMan{})       //	Uncomment this line for Pacman integration
	mgr.AddManager("winget", Manager.WingetMan{})    //	Uncomment this line for Winget integration
	mgr.AddManager("scoop", Manager.ScoopMan{})      //	Uncomment this line for Scoop integration
	mgr.AddManager("chocolatey", Manager.ChocoMan{}) //	Uncomment this line for Chocolatey integration
	mgr.AddManager("npm", Manager.NpmMan{})          //	Uncomment this line for NPM integration
	mgr.AddManager("nix", Manager.NixMan{})			//	Uncomment this line for Nix Integration

	var usern string
	fmt.Print("Please enter username:")
	fmt.Scanln(&usern)
	for len(usern) > 0 {
		fmt.Println("Press 'L' for list, 'I' for install, and 'R' for remove")
		var choice string
		fmt.Print("> ")
		fmt.Scanln(&choice)

		if choice == "L" || choice == "l" {
			mgr.ListPackages()
		}
		if choice == "i" || choice == "I" {
			fmt.Print("Please enter a package name: ")
			var whichpkg string
			fmt.Scanln(&whichpkg)
			mgr.InstallPackage(whichpkg)
		}
		if choice == "r" || choice == "R" {
			fmt.Print("Please enter a package name: ")
			var whichpkg string
			fmt.Scanln(&whichpkg)
			mgr.RemovePackage(whichpkg)
		}
	}
}
