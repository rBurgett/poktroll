package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/pokt-network/poktroll/telemetry"
	"github.com/pokt-network/poktroll/x/gateway/types"
)

// TODO_BLOCKER(#489): Apps & gateways unbonding periods
func (k msgServer) UnstakeGateway(
	goCtx context.Context,
	msg *types.MsgUnstakeGateway,
) (*types.MsgUnstakeGatewayResponse, error) {
	isSuccessful := false
	defer telemetry.EventSuccessCounter(
		"unstake_gateway",
		telemetry.DefaultCounterFn,
		func() bool { return isSuccessful },
	)

	ctx := sdk.UnwrapSDKContext(goCtx)

	logger := k.Logger().With("method", "UnstakeGateway")
	logger.Info(fmt.Sprintf("About to unstake gateway with msg: %v", msg))

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Check if the gateway already exists or not
	var err error
	gateway, isGatewayFound := k.GetGateway(ctx, msg.Address)
	if !isGatewayFound {
		logger.Info(fmt.Sprintf("Gateway not found. Cannot unstake address %s", msg.Address))
		return nil, types.ErrGatewayNotFound
	}
	logger.Info(fmt.Sprintf("Gateway found. Unstaking gateway for address %s", msg.Address))

	// Retrieve the address of the gateway
	gatewayAddress, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		logger.Error(fmt.Sprintf("could not parse address %s", msg.Address))
		return nil, err
	}

	// Send the coins from the gateway pool back to the gateway
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, gatewayAddress, []sdk.Coin{*gateway.Stake})
	if err != nil {
		logger.Error(fmt.Sprintf("could not send %v coins from %s module to %s account due to %v", gateway.Stake, gatewayAddress, types.ModuleName, err))
		return nil, err
	}

	// Update the Gateway in the store
	k.RemoveGateway(ctx, gatewayAddress.String())
	logger.Info(fmt.Sprintf("Successfully removed the gateway: %+v", gateway))

	err = ctx.EventManager().EmitTypedEvent(&types.EventGatewayUnstaked{Address: msg.Address})
	if err != nil {
		logger.Error(fmt.Sprintf("failed to emit event for unstaking gateway: %v", err))
		return nil, err
	}

	isSuccessful = true
	return &types.MsgUnstakeGatewayResponse{}, nil
}
