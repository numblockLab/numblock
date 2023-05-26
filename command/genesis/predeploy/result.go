package predeploy

import (
	"bytes"
	"fmt"

	"https://github.com/numblockLab/numblock/command/helper"
)

type GenesisPredeployResult struct {
	Address string `json:"address"`
}

func (r *GenesisPredeployResult) GetOutput() string {
	var buffer bytes.Buffer

	buffer.WriteString("\n[SMART CONTRACT PREDEPLOYMENT]\n")

	outputs := []string{
		fmt.Sprintf("Address|%s", r.Address),
	}

	buffer.WriteString(helper.FormatKV(outputs))
	buffer.WriteString("\n")

	return buffer.String()
}