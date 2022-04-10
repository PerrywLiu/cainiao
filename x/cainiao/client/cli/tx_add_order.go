package cli

import (
	"strconv"

	"github.com/cosmonaut/cainiao/x/cainiao/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-order [send-name] [send-address] [send-tel] [receive-name] [receive-address] [receive-tel]",
		Short: "Broadcast message add-order",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSendName := args[0]
			argSendAddress := args[1]
			argSendTel := args[2]
			argReceiveName := args[3]
			argReceiveAddress := args[4]
			argReceiveTel := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddOrder(
				clientCtx.GetFromAddress().String(),
				argSendName,
				argSendAddress,
				argSendTel,
				argReceiveName,
				argReceiveAddress,
				argReceiveTel,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
