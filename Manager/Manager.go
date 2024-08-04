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

func (m *ManagerList) InstallPackage(pack string) {
	for name, manager := range m.managers {
		installed, err := manager.IsInstalled(pack)
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

type BrewMan struct{}

func (s BrewMan) PkgListInstalled() ([]string, error) {
	cmd := exec.Command("brew", "list")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

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

func (s BrewMan) PkgInstall(pack string) ([]string, error) {
	cmd := exec.Command("brew", "install", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

type SnapMan struct{}

func (s SnapMan) PkgListInstalled() ([]string, error) {
	cmd := exec.Command("snap", "list")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

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

func (s SnapMan) PkgInstall(pack string) ([]string, error) {
	cmd := exec.Command("snap", "install", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}
