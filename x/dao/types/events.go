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

package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	EventTypeCreateVoter = "create_voter"
	AttributeKeyVoterID  = "voter_id"
	AttributeKeyCreator  = "creator"
	AttributeKeyWeight   = "weight"

	EventTypeVoteProposal   = "vote_proposal"
	AttributeKeyProposalID  = "proposal_id"
	AttributeKeyDecision    = "decision"
	AttributeKeyExplanation = "explanation"
)

func NewEventCreateVoter(id uint64, creator sdk.AccAddress, weight uint64) sdk.Event {
	return sdk.NewEvent(
		EventTypeCreateVoter,
		sdk.NewAttribute(AttributeKeyVoterID, fmt.Sprintf("%d", id)),
		sdk.NewAttribute(AttributeKeyCreator, creator.String()),
		sdk.NewAttribute(AttributeKeyWeight, fmt.Sprintf("%d", weight)),
	)
}

func NewEventVoteProposal(proposalID uint64, voterID uint64, decision VoteType, explanation string) sdk.Event {
	return sdk.NewEvent(
		EventTypeVoteProposal,
		sdk.NewAttribute(AttributeKeyProposalID, fmt.Sprintf("%d", proposalID)),
		sdk.NewAttribute(AttributeKeyVoterID, fmt.Sprintf("%d", voterID)),
		sdk.NewAttribute(AttributeKeyDecision, decision.String()),
		sdk.NewAttribute(AttributeKeyExplanation, explanation),
	)
}
