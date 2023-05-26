package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/numblockLab/numblock/command/backup"
	"github.com/numblockLab/numblock/command/bridge"
	"github.com/numblockLab/numblock/command/genesis"
	"github.com/numblockLab/numblock/command/helper"
	"github.com/numblockLab/numblock/command/ibft"
	"github.com/numblockLab/numblock/command/license"
	"github.com/numblockLab/numblock/command/monitor"
	"github.com/numblockLab/numblock/command/peers"
	"github.com/numblockLab/numblock/command/polybft"
	"github.com/numblockLab/numblock/command/polybftmanifest"
	"github.com/numblockLab/numblock/command/polybftsecrets"
	"github.com/numblockLab/numblock/command/regenesis"
	"github.com/numblockLab/numblock/command/rootchain"
	"github.com/numblockLab/numblock/command/secrets"
	"github.com/numblockLab/numblock/command/server"
	"github.com/numblockLab/numblock/command/status"
	"github.com/numblockLab/numblock/command/txpool"
	"github.com/numblockLab/numblock/command/version"
	"github.com/numblockLab/numblock/command/whitelist"
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
