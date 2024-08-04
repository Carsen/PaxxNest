package Manager

type PackageManager interface {
	ListInstalledPackages() ([]string, error)
	InstallPackage(name string) error
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
