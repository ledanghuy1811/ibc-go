package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller/types"
	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
)

var _ types.QueryServer = Keeper{}

// InterchainAccount implements the Query/InterchainAccount gRPC method
func (k Keeper) InterchainAccount(goCtx context.Context, req *types.QueryInterchainAccountRequest) (*types.QueryInterchainAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	portID, err := icatypes.NewControllerPortID(req.Owner)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to generate portID from owner address: %s", err)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	addr, found := k.GetInterchainAccountAddress(ctx, req.ConnectionId, portID)
	if !found {
		return nil, status.Errorf(codes.NotFound, "failed to retrieve account address for %s on connection %s", portID, req.ConnectionId)
	}

	return &types.QueryInterchainAccountResponse{
		Address: addr,
	}, nil
}

// Params implements the Query/Params gRPC method
func (k Keeper) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)

	chanel, typeChanle := k.channelKeeper.GetChannel(ctx, "icacontroller-orai10lv5lcj22zfnfqrqyqq73w5nvkjvmqa9a277lhc0vyu2hpm9ngssenvaaq", "channel-55")
	ctx.Logger().Error(fmt.Sprintf("chanel: %v and type: %v", chanel, typeChanle))

	chanCap, f := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath("icacontroller-orai10lv5lcj22zfnfqrqyqq73w5nvkjvmqa9a277lhc0vyu2hpm9ngssenvaaq", "channel-55"))
	ctx.Logger().Error(fmt.Sprintf("v4 capability query here: %v and found: %v", chanCap, f))

	capa, found := k.scopedKeeper.GetCapability(ctx, "capabilities/ports/icacontroller-orai10lv5lcj22zfnfqrqyqq73w5nvkjvmqa9a277lhc0vyu2hpm9ngssenvaaq/channels/channel-55")
	ctx.Logger().Error(fmt.Sprintf("v8 capability query here: %v and found: %v", capa, found))

	return &types.QueryParamsResponse{
		Params: &params,
	}, nil
}
