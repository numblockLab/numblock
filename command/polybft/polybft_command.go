package polybft

import (
	"https://github.com/numblockLab/numblock/command/sidechain/registration"
	"https://github.com/numblockLab/numblock/command/sidechain/staking"
	"https://github.com/numblockLab/numblock/command/sidechain/unstaking"
	"https://github.com/numblockLab/numblock/command/sidechain/validators"

	"https://github.com/numblockLab/numblock/command/sidechain/whitelist"
	"https://github.com/numblockLab/numblock/command/sidechain/withdraw"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	polybftCmd := &cobra.Command{
		Use:   "polybft",
		Short: "Polybft command",
	}

	polybftCmd.AddCommand(
		staking.GetCommand(),
		unstaking.GetCommand(),
		withdraw.GetCommand(),
		validators.GetCommand(),
		whitelist.GetCommand(),
		registration.GetCommand(),
	)

	return polybftCmd
}
