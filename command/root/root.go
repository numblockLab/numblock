package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"https://github.com/numblockLab/numblock/command/backup"
	"https://github.com/numblockLab/numblock/command/bridge"
	"https://github.com/numblockLab/numblock/command/genesis"
	"https://github.com/numblockLab/numblock/command/helper"
	"https://github.com/numblockLab/numblock/command/ibft"
	"https://github.com/numblockLab/numblock/command/license"
	"https://github.com/numblockLab/numblock/command/monitor"
	"https://github.com/numblockLab/numblock/command/peers"
	"https://github.com/numblockLab/numblock/command/polybft"
	"https://github.com/numblockLab/numblock/command/polybftmanifest"
	"https://github.com/numblockLab/numblock/command/polybftsecrets"
	"https://github.com/numblockLab/numblock/command/regenesis"
	"https://github.com/numblockLab/numblock/command/rootchain"
	"https://github.com/numblockLab/numblock/command/secrets"
	"https://github.com/numblockLab/numblock/command/server"
	"https://github.com/numblockLab/numblock/command/status"
	"https://github.com/numblockLab/numblock/command/txpool"
	"https://github.com/numblockLab/numblock/command/version"
	"https://github.com/numblockLab/numblock/command/whitelist"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Polygon Edge is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		rootchain.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		whitelist.GetCommand(),
		license.GetCommand(),
		polybftsecrets.GetCommand(),
		polybft.GetCommand(),
		polybftmanifest.GetCommand(),
		bridge.GetCommand(),
		regenesis.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
