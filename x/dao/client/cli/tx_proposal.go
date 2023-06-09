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
	"errors"
	"regexp"
	"strconv"
	"time"

	"lordverse/x/dao/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-proposal [warehouse-id] [title] [description] [expiration]",
		Short: "Create a new proposal",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			warehouseID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argTitle := args[1]
			argDescription := args[2]
			argExpiration, err := parseDuration(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateProposal(clientCtx.GetFromAddress().String(), warehouseID, argTitle, argDescription, &argExpiration)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-proposal [id] [description]",
		Short: "Update a proposal",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argDescription := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateProposal(clientCtx.GetFromAddress().String(), id, argDescription)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// func CmdVoteProposal() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "vote-proposal [id] [voter-id] [decision]",
// 		Short: "Vote a proposal",
// 		Args:  cobra.ExactArgs(3),
// 		RunE: func(cmd *cobra.Command, args []string) (err error) {
// // 			argProposalId, err := strconv.ParseUint(args[0], 10, 64)
// // 			if err != nil {
// // 				return err
// // 			}

// // 			argDecision, err := strconv.ParseUint(args[1], 10, 64)
// // 			if err != nil {
// // 				return err
// // 			}
// // vote := types.VoteType(argDecision)

// // clientCtx, err := client.GetClientTxContext(cmd)
// // if err != nil {
// // 	return err
// // }

// // msg := types.NewMsgVoteProposal(clientCtx.GetFromAddress().String(), argProposalId, vote)
// // if err := msg.ValidateBasic(); err != nil {
// // 	return err
// // }
// // return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
// return nil,nil

// 	}
// }
// 	flags.AddTxFlagsToCmd(cmd)

// 	return cmd
// }

func CmdVoteProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vote-proposal [proposal-id] [voter-id] [decision] [explanation]",
		Short: "vote a proposal by id",
		Long:  "vote a proposal by id. decision can be 0,1,2 which is abstain, yes, no",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			proposalID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			voterID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}
			// 0: abstain 1: yes 2: no
			decision, err := strconv.ParseInt(args[2], 10, 32)
			if err != nil {
				return err
			}

			explanation := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgVoteProposal(clientCtx.GetFromAddress().String(), proposalID, voterID, types.VoteType(decision), explanation)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-proposal [id]",
		Short: "Delete a proposal by id",
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

			msg := types.NewMsgDeleteProposal(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// parseDuration parses a duration string and returns a future time.
// The duration string should be in the format of "xdaysxhoursxminutesxseconds".
// Code generated by Codeium with prompt template:
// Code generation
// I have a time in format 1days2hours3minutes4seconds which means 1 day 2 hours 3 minutes 4 seconds start from now, I want to convert it to a time.Time
func parseDuration(s string) (time.Time, error) {
	// Use regular expressions to extract the duration components
	r := regexp.MustCompile(`(\d+)days(\d+)hours(\d+)minutes(\d+)seconds`)
	matches := r.FindStringSubmatch(s)
	if len(matches) != 5 {
		return time.Time{}, errors.New("invalid duration format,should be in the format of \"xdaysxhoursxminutesxseconds\"")
	}

	// Parse the duration components into integers
	days, _ := strconv.Atoi(matches[1])
	hours, _ := strconv.Atoi(matches[2])
	minutes, _ := strconv.Atoi(matches[3])
	seconds, _ := strconv.Atoi(matches[4])

	// Calculate the total duration in seconds
	totalSeconds := (days * 24 * 60 * 60) + (hours * 60 * 60) + (minutes * 60) + seconds

	// Get the current time and add the duration to it
	now := time.Now()
	duration := time.Duration(totalSeconds) * time.Second
	future := now.Add(duration)

	// Return the future time
	return future, nil
}
