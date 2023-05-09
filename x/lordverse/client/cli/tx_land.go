package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"lordverse/x/lordverse/types"
)

func CmdCreateLand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-land [owner] [location] [grade] [assessment] [extra]",
		Short: "Create a new Land",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOwner := args[0]
			argLocation := args[1]
			argGrade := args[2]
			argAssessment := args[3]
			argExtra := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateLand(clientCtx.GetFromAddress().String(), argOwner, argLocation, argGrade, argAssessment, argExtra)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateLand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-land [id] [owner] [location] [grade] [assessment] [extra]",
		Short: "Update a Land",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argOwner := args[1]

			argLocation := args[2]

			argGrade := args[3]

			argAssessment := args[4]

			argExtra := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateLand(clientCtx.GetFromAddress().String(), id, argOwner, argLocation, argGrade, argAssessment, argExtra)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteLand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-land [id]",
		Short: "Delete a Land by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteLand(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
