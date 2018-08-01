package node

import (
	"fmt"

	"github.com/spf13/cobra"
)

var chaincodeDevMode bool

func startCmd() *cobra.Command {
	return nodeStartCmd
}

var nodeStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the node.",
	Long:  `Starts a node that interacts with the network.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return serve(args)
	},
}

func serve(args []string) error {
	fmt.Println("不要企图爱上哥，哥只是个传说:", args)
	return nil
}
