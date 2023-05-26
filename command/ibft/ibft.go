package ibft

import (
	"https://github.com/numblockLab/numblock/command/helper"
	"https://github.com/numblockLab/numblock/command/ibft/candidates"
	"https://github.com/numblockLab/numblock/command/ibft/propose"
	"https://github.com/numblockLab/numblock/command/ibft/quorum"
	"https://github.com/numblockLab/numblock/command/ibft/snapshot"
	"https://github.com/numblockLab/numblock/command/ibft/status"
	_switch "https://github.com/numblockLab/numblock/command/ibft/switch"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	ibftCmd := &cobra.Command{
		Use:   "ibft",
		Short: "Top level IBFT command for interacting with the IBFT consensus. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(ibftCmd)

	registerSubcommands(ibftCmd)

	return ibftCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// ibft status
		status.GetCommand(),
		// ibft snapshot
		snapshot.GetCommand(),
		// ibft propose
		propose.GetCommand(),
		// ibft candidates
		candidates.GetCommand(),
		// ibft switch
		_switch.GetCommand(),
		// ibft quorum
		quorum.GetCommand(),
	)
}
