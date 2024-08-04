package Manager

import (
	"fmt"
	"log"
	"os/exec"
)

// Interface to define common Package Manager Operations
type PkgMgrOps interface {
	PkgListInstalled() ([]string, error)
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
		fmt.Printf("Installing %s", pack, "via manager %s\n", name)
		output, err := manager.PkgInstall(pack)
		if err != nil {
			log.Fatal(err)
			continue
		}
		for _, pkg := range packages {
			fmt.Println(pkg)
		}
	}
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

func (s BrewMan) PkgInstall(name string) ([]string, error) {
	cmd := exec.Command("brew", "install", name)
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

func (s SnapMan) PkgInstall(name string) ([]string, error) {
	cmd := exec.Command("snap", "install", name)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}
