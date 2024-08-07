package Manager

import (
	"fmt"
	"os/exec"

	"github.com/Carsen/PaxxNest/ErrLog"
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

func NewManager() *ManagerList {
	return &ManagerList{
		managers: make(map[string]PkgMgrOps),
	}
}

func (m *ManagerList) AddManager(name string, manager PkgMgrOps) {
	m.managers[name] = manager
}

// List Packages
func (m *ManagerList) ListPackages() {
	for name, manager := range m.managers {
		packages, err := manager.PkgListInstalled()
		if err != nil {
			ErrLog.LogErr(err)
			continue
		}
		for _, pkg := range packages {
			fmt.Printf("Packages from %s:\n", name)
			fmt.Println(pkg)
		}
	}
}

// Install Packages
func (m *ManagerList) InstallPackage(pack string) {
	for name, manager := range m.managers {
		installed, err := manager.PkgIsInstalled(pack)
		if err != nil {
			ErrLog.LogErr(err)
		}
		fmt.Printf("Trying to install '%s' via manager '%s'.\n", pack, name)
		if installed {
			fmt.Printf("Package '%s' is already installed via manager '%s'.\n", pack, name)
			return
		}

		output, err := manager.PkgInstall(pack)
		if err != nil {
			fmt.Printf("Failed to install '%s' using manager '%s': %v. \nTrying Next...\n", pack, name, err)
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
		fmt.Printf("Trying to remove '%s', via manager '%s'.\n", pack, name)
		if err != nil {
			ErrLog.LogErr(err)
			continue
		}
		if !installed {
			fmt.Printf("Package '%s' is not currently installed.", pack)
			return
		}
		output, err := manager.PkgRemove(pack)
		if err != nil {
			fmt.Printf("Failed to remove '%s' using manager '%s': %v.\n Trying next", pack, name, err)
			continue
		}

		fmt.Printf("%s: \n%s\n", name, output)
		fmt.Printf("Successfully removed '%s' using manager '%s'.\n", pack, name)
		return
	}
	fmt.Printf("Failed to remove '%s' using all available package managers. \n", pack)
}

// Integration with HomeBrew Package Manager
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
	_, err := cmd.CombinedOutput()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s BrewMan) PkgInstall(pack string) ([]string, error) {
	cmd := exec.Command("brew", "install", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

func (s BrewMan) PkgRemove(pack string) ([]string, error) {
	cmd := exec.Command("brew", "remove", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

// Integration with Snap Package Manager
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
	_, err := cmd.CombinedOutput()
	if err != nil {
		return false, err
	}
	return true, err
}

func (s SnapMan) PkgInstall(pack string) ([]string, error) {
	cmd := exec.Command("snap", "install", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

func (s SnapMan) PkgRemove(pack string) ([]string, error) {
	cmd := exec.Command("snap", "remove", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

// Integration with PacMan Package Manager
type PacMan struct{}

func (s PacMan) PkgListInstalled() ([]string, error) {
	cmd := exec.Command("pacman", "-Q")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

func (s PacMan) PkgIsInstalled(pack string) (bool, error) {
	cmd := exec.Command("pacman", "-Q", pack)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s PacMan) PkgInstall(pack string) ([]string, error) {
	cmd := exec.Command("pacman", "-S", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

func (s PacMan) PkgRemove(pack string) ([]string, error) {
	cmd := exec.Command("pacman", "-R", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

// Integration with Winget Manager
type WingetMan struct{}

func (s WingetMan) PkgListInstalled() ([]string, error) {
	cmd := exec.Command("winget", "list")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

func (s WingetMan) PkgIsInstalled(pack string) (bool, error) {
	cmd := exec.Command("winget", "list", pack)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s WingetMan) PkgInstall(pack string) ([]string, error) {
	cmd := exec.Command("winget", "install", pack, "--accept-source-agreements --accept-package-agreements")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

func (s WingetMan) PkgRemove(pack string) ([]string, error) {
	cmd := exec.Command("winget", "remove", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

// Integration with Scoop Manager
type ScoopMan struct{}

func (s ScoopMan) PkgListInstalled() ([]string, error) {
	cmd := exec.Command("scoop", "list")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

func (s ScoopMan) PkgIsInstalled(pack string) (bool, error) {
	cmd := exec.Command("scoop", "list", pack)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false, err
	}
	return len(string(output)) > 50, nil
}

func (s ScoopMan) PkgInstall(pack string) ([]string, error) {
	cmd := exec.Command("scoop", "install", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

func (s ScoopMan) PkgRemove(pack string) ([]string, error) {
	cmd := exec.Command("scoop", "uninstall", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

// Integration with Chocolatey manager
type ChocoMan struct{}

func (s ChocoMan) PkgListInstalled() ([]string, error) {
	cmd := exec.Command("choco", "list", "--local-only")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

func (s ChocoMan) PkgIsInstalled(pack string) (bool, error) {
	cmd := exec.Command("choco", "list", pack, "--local-only")
	_, err := cmd.Output()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s ChocoMan) PkgInstall(pack string) ([]string, error) {
	cmd := exec.Command("choco", "install", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

func (s ChocoMan) PkgRemove(pack string) ([]string, error) {
	cmd := exec.Command("choco", "uninstall", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

// Integration with NPM manager
type NpmMan struct{}

func (s NpmMan) PkgListInstalled() ([]string, error) {
	cmd := exec.Command("npm", "list")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

func (s NpmMan) PkgIsInstalled(pack string) (bool, error) {
	cmd := exec.Command("npm", "list", pack)
	_, err := cmd.Output()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s NpmMan) PkgInstall(pack string) ([]string, error) {
	cmd := exec.Command("npm", "install", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

func (s NpmMan) PkgRemove(pack string) ([]string, error) {
	cmd := exec.Command("npm", "uninstall", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

// Integration with NIX Manager
type NixMan struct{}

func (s NixMan) PkgListInstalled() ([]string, error) {
	cmd := exec.Command("nix-env --query --installed")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

func (s NixMan) PkgIsInstalled(pack string) (bool, error) {
	cmd := exec.Command("nix-env -qaP", pack)
	_, err := cmd.Output()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s NixMan) PkgInstall(pack string) ([]string, error) {
	command := fmt.Sprintf("nix-env -iA nixpkgs.%s", pack)
	cmd := exec.Command(command)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}

func (s NixMan) PkgRemove(pack string) ([]string, error) {
	cmd := exec.Command("nix-env -e", pack)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return []string{string(output)}, nil
}
