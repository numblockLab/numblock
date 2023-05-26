package main

import (
	_ "embed"

	"https://github.com/numblockLab/numblock/command/root"
	"https://github.com/numblockLab/numblock/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
