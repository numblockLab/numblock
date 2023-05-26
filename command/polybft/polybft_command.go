package polybft

import (
	"github.com/numblockLab/numblock/command/sidechain/registration"
	"github.com/numblockLab/numblock/command/sidechain/staking"
	"github.com/numblockLab/numblock/command/sidechain/unstaking"
	"github.com/numblockLab/numblock/command/sidechain/validators"

	"github.com/numblockLab/numblock/command/sidechain/whitelist"
	"github.com/numblockLab/numblock/command/sidechain/withdraw"
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
