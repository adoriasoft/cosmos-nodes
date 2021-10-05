package cosmosnodes_test

import (
	"testing"

	keepertest "github.com/adoriasoft/cosmos-nodes/testutil/keeper"
	"github.com/adoriasoft/cosmos-nodes/x/cosmosnodes"
	"github.com/adoriasoft/cosmos-nodes/x/cosmosnodes/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmosnodesKeeper(t)
	cosmosnodes.InitGenesis(ctx, *k, genesisState)
	got := cosmosnodes.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
