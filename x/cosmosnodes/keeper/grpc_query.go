package keeper

import (
	"github.com/adoriasoft/cosmos-nodes/x/cosmosnodes/types"
)

var _ types.QueryServer = Keeper{}
