package status

import (
	"context"

	"https://github.com/numblockLab/numblock/command"
	"https://github.com/numblockLab/numblock/command/helper"
	ibftOp "https://github.com/numblockLab/numblock/consensus/ibft/proto"
	"github.com/spf13/cobra"
	empty "google.golang.org/protobuf/types/known/emptypb"
)

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Returns the current validator key of the IBFT client",
		Run:   runCommand,
	}
}

func runCommand(cmd *cobra.Command, _ []string) {
	outputter := command.InitializeOutputter(cmd)
	defer outputter.WriteOutput()

	statusResponse, err := getIBFTStatus(helper.GetGRPCAddress(cmd))
	if err != nil {
		outputter.SetError(err)

		return
	}

	outputter.SetCommandResult(&IBFTStatusResult{
		ValidatorKey: statusResponse.Key,
	})
}

func getIBFTStatus(grpcAddress string) (*ibftOp.IbftStatusResp, error) {
	client, err := helper.GetIBFTOperatorClientConnection(
		grpcAddress,
	)
	if err != nil {
		return nil, err
	}

	return client.Status(context.Background(), &empty.Empty{})
}
