package Manager

import (
	"fmt"
	"log"
	"os/exec"
)

// Interface to define common Package Manager Operations
type PkgMgrOps interface {
	PkgListInstalled() ([]string, error)
	PkgIsInstalled(pack string) (bool, error)
	PkgInstall(name string) ([]string, error)
	PkgRemove(name string) ([]string, error)
}

// Struct for different package managers
type ManagerList struct {
	managers map[string]PkgMgrOps
}

// Create new instance of ManagerList struct
func NewManager() *ManagerList {
	return &ManagerList{
		managers: make(map[string]PkgMgrOps),
	}
}

// Add a manager to ManagerList
func (m *ManagerList) AddManager(name string, manager PkgMgrOps) {
	m.managers[name] = manager
}

// List Packages
func (m *ManagerList) ListPackages() {
	for name, manager := range m.managers {
		fmt.Printf("Packages from %s:\n", name)
		packages, err := manager.PkgListInstalled()
		if err != nil {
			log.Fatal(err)
			continue
		}
		for _, pkg := range packages {
			fmt.Println(pkg)
		}
	}
}

// Install Packages
func (m *ManagerList) InstallPackage(pack string) {
	for name, manager := range m.managers {
		installed, err := manager.PkgIsInstalled(pack)
		if err != nil {
			log.Fatalf("Error checking if package '%s' is installed: %v", pack, err)
		}

		if installed {
			fmt.Printf("Package '%s' is already installed via manager '%s'.\n", pack, name)
			return
		}

		fmt.Printf("Installing '%s', via manager '%s'\n", pack, name)
		output, err := manager.PkgInstall(pack)
		if err != nil {
			fmt.Printf("Failed to install '%s' using manager '%s': %v. \n Trying Next", pack, name, err)
			continue
		}

		fmt.Printf("%s: \n%s\n", pack, output)
		fmt.Printf("Successfully installed '%s' via manager '%s'.\n", pack, name)
		return
	}
	fmt.Printf("Failed to install '%s' using all available packages managers.\n", pack)
}

// Remove Packages
func (m *ManagerList) RemovePackage(pack string) {
	for name, manager := range m.managers {
		installed, err := manager.PkgIsInstalled(pack)
		if err != nil {
			log.Fatalf("Error checking if package '%s' is installed: %v", pack, err)
		}

		if !installed {
			fmt.Printf("Package '%s' is not currently installed.", pack)
			return
		}

		fmt.Printf("Removing '%s', via manager '%s'.\n", pack, name)
		output, err := manager.PkgInstall(pack)
		if err != nil {
			fmt.Printf("Failed to remove '%s' using manager '%s': %v.\n Trying next", pack, name, err)
			continue
		}

		fmt.Printf("%s: \n%s\n", pack, output)
		fmt.Printf("Successfully removed '%s' using manager '%s'.\n", pack, name)
		return
	}
	fmt.Printf("Failed to remove '%s' using all available package managers. \n", pack)
}

//
//
//
// Integration with HomeBrew Package Manager
type BrewMan struct{}

// Brew List
func (s BrewMan) PkgListInstalled() ([]string, error) {
	cmd := exec.Command("brew", "list")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

// Check if 'Package' was installed using Brew List
func (s BrewMan) PkgIsInstalled(pack string) (bool, error) {
	cmd := exec.Command("brew", "list", pack)
	output, err := cmd.CombinedOutput()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
			return false, nil
		}
	}
	return len(output) > 0, nil
}

// Brew Install 'Package'
func (s BrewMan) PkgInstall(pack string) ([]string, error) {
	cmd := exec.Command("brew", "install", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

// Brew Remove 'Package'
func (s BrewMan) PkgRemove(pack string) ([]string, error) {
	cmd := exec.Command("brew", "remove", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

//
//
//
// Integration with Snap Package Manager
type SnapMan struct{}

// Snap List
func (s SnapMan) PkgListInstalled() ([]string, error) {
	cmd := exec.Command("snap", "list")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

// Check is 'Package' is installed using Snap List
func (s SnapMan) PkgIsInstalled(pack string) (bool, error) {
	cmd := exec.Command("snap", "list", pack)
	output, err := cmd.CombinedOutput()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
			return false, nil
		}
	}
	return len(output) > 0, nil
}

// Snap Install 'Package'
func (s SnapMan) PkgInstall(pack string) ([]string, error) {
	cmd := exec.Command("snap", "install", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

// Snap Remove 'Package'
func (s SnapMan) PkgRemove(pack string) ([]string, error) {
	cmd := exec.Command("snap", "remove", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

//
//
//
// Integration with PacMan Package Manager
type PacMan struct{}

// PacMan List
func (s PacMan) PkgListInstalled() ([]string, error) {
	cmd := exec.Command("pacman", "-Q")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

// Check is 'Package' is installed using Pacman List
func (s PacMan) PkgIsInstalled(pack string) (bool, error) {
	cmd := exec.Command("pacman", "-Q", pack)
	output, err := cmd.CombinedOutput()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
			return false, nil
		}
	}
	return len(output) > 0, nil
}

// Pacman Install 'Package'
func (s PacMan) PkgInstall(pack string) ([]string, error) {
	cmd := exec.Command("pacman", "-S", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

// Pacman Remove 'Package'
func (s PacMan) PkgRemove(pack string) ([]string, error) {
	cmd := exec.Command("pacman", "-R", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

//
//
//
// Integration with Winget Manager
type WingetMan struct{}

// Winget List
func (s WingetMan) PkgListInstalled() ([]string, error) {
	cmd := exec.Command("winget", "list")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

// Check is 'Package' is installed using Winget List
func (s WingetMan) PkgIsInstalled(pack string) (bool, error) {
	cmd := exec.Command("winget", "list", pack)
	output, err := cmd.CombinedOutput()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
			return false, nil
		}
	}
	return len(output) > 0, nil
}

// Winget Install 'Package'
func (s WingetMan) PkgInstall(pack string) ([]string, error) {
	cmd := exec.Command("winget", "install", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

// Winget Remove 'Package'
func (s WingetMan) PkgRemove(pack string) ([]string, error) {
	cmd := exec.Command("winget", "remove", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

//
//
//
// Integration with Scoop Manager
type ScoopMan struct{}

// Scoop List
func (s ScoopMan) PkgListInstalled() ([]string, error) {
	cmd := exec.Command("scoop", "list")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

// Check is 'Package' is installed using Scoop List
func (s ScoopMan) PkgIsInstalled(pack string) (bool, error) {
	cmd := exec.Command("scoop", "list", pack)
	output, err := cmd.CombinedOutput()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
			return false, nil
		}
	}
	return len(output) > 0, nil
}

// Scoop Install 'Package'
func (s ScoopMan) PkgInstall(pack string) ([]string, error) {
	cmd := exec.Command("scoop", "install", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

// Scoop Remove 'Package'
func (s ScoopMan) PkgRemove(pack string) ([]string, error) {
	cmd := exec.Command("scoop", "uninstall", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

//
//
//
// Integration with Chocolatey manager
type ChocoMan struct{}

// Choco List
func (s ChocoMan) PkgListInstalled() ([]string, error) {
	cmd := exec.Command("choco", "list", "--local-only")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}


// Check if 'Package' is installed using Choco List
func (s ChocoMan) PkgIsInstalled(pack string) (bool, error) {
	cmd := exec.Command("choco", "list", pack, "--local-only")
	output, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
			return false, nil
		}
	}
	return len(output) > 0, nil
}

// Choco install 'Package'
func (s ChocoMan) PkgInstall(pack string) ([]string, error) {
	cmd := exec.Command("choco", "install", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

// Choco remove 'Package'
func (s ChocoMan) PkgRemove(pack string) ([]string, error) {
	cmd := exec.Command("choco", "uninstall", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

//
//
//
// Integration with NPM manager
type NpmMan struct{}

// NPM List
func (s NpmMan) PkgListInstalled() ([]string, error) {
	cmd := exec.Command("npm", "list")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}


// Check if 'Package' is installed using NPM List
func (s NpmMan) PkgIsInstalled(pack string) (bool, error) {
	cmd := exec.Command("npm", "list", pack)
	output, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
			return false, nil
		}
	}
	return len(output) > 0, nil
}

// NPM install 'Package'
func (s NpmMan) PkgInstall(pack string) ([]string, error) {
	cmd := exec.Command("npm", "install", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

// NPM remove 'Package'
func (s NpmMan) PkgRemove(pack string) ([]string, error) {
	cmd := exec.Command("npm", "uninstall", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}
