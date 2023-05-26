package main

import (
	_ "embed"

	"github.com/numblockLab/numblock/command/root"
	"github.com/numblockLab/numblock/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
