package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"pocket/x/supplier/types"
)

func (k msgServer) CreateClaim(goCtx context.Context, msg *types.MsgCreateClaim) (*types.MsgCreateClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	/*
		INCOMPLETE: Handling the message

		## Validation

		### Session validation
		1. [ ] claimed session ID == retrieved session ID
		2. [ ] this supplier is in the session's suppliers list

		### Msg distribution validation (depends on session validation)
		1. [ ] pseudo-randomize earliest block offset
		2. [ ] governance-based earliest block offset

		### Claim validation
		1. [ ] session validation
		2. [ ] msg distribution validation

		## Persistence
		1. [ ] create claim message
			- supplier address
			- session header
			- claim
		2. [ ] last block height commitment; derives:
			- last block committed hash, must match proof path
			- session ID (?)
	*/
	_ = ctx

	return &types.MsgCreateClaimResponse{}, nil
}