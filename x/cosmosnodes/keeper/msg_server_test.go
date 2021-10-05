package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/adoriasoft/cosmos-nodes/testutil/keeper"
	"github.com/adoriasoft/cosmos-nodes/x/cosmosnodes/keeper"
	"github.com/adoriasoft/cosmos-nodes/x/cosmosnodes/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmosnodesKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
