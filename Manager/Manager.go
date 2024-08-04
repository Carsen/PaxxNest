package Manager

import (
	"fmt"
	"log"
	"os/exec"
)

// Interface to define common Package Manager Operations
type PkgMgrOps interface {
	ListInstalledPackages() ([]string, error)
	InstallPackage(name string) error
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
		packages, err := manager.ListInstalledPackages()
		if err != nil {
			log.Fatal(err)
			continue
		}
		for _, pkg := range packages {
			fmt.Println(pkg)
		}
	}
}

type SnapMan struct{}

func (s SnapMan) ListInstalledPackages() ([]string, error) {
	cmd := exec.Command("snap", "list")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}
