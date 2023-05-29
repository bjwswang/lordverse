/*
Copyright 2023 The Bestchains Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cli

import (
	"strconv"
	"strings"

	"lordverse/x/dao/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateWarehouse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-warehouse [voters] [threshold]",
		Short: "Create a new warehouse",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argVoters := parseArgVoters(args[0])

			argThreshold, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateWarehouse(clientCtx.GetFromAddress().String(), argVoters, argThreshold)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdSignWarehouse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign-warehouse [warehouse-id] [voter-id] ",
		Short: "Sign a warehouse",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			warehouseID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			voterID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSignWarehouse(clientCtx.GetFromAddress().String(), warehouseID, voterID)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateWarehouse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-warehouse [id] [voters] [threshold]",
		Short: "Update a warehouse",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argVoters := args[1]
			// convert argVoters to string array seperated by commas
			argVotersArray := parseArgVoters(argVoters)

			argThreshold, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateWarehouse(clientCtx.GetFromAddress().String(), id, argVotersArray, argThreshold)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteWarehouse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-warehouse [id]",
		Short: "Delete a warehouse by id",
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

			msg := types.NewMsgDeleteWarehouse(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// parseArgVoters parses a comma-separated string of voters into a slice of uint64 values.
// If a value fails to parse, it is skipped and the next value is attempted.
// Returns a slice of uint64 values.
func parseArgVoters(voters string) []uint64 {
	argVotersArray := strings.Split(voters, ",")
	argVotersArrayUint64 := make([]uint64, len(argVotersArray))
	for i, voter := range argVotersArray {
		argVotersArrayUint64[i], _ = strconv.ParseUint(voter, 10, 64)
	}
	return argVotersArrayUint64
}
