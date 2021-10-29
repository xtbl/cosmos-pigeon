package keeper

import (
	"github.com/cosmonaut/pigeon/x/pigeon/types"
)

var _ types.QueryServer = Keeper{}
