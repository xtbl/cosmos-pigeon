package pigeon_test

import (
	"testing"

	keepertest "github.com/cosmonaut/pigeon/testutil/keeper"
	"github.com/cosmonaut/pigeon/x/pigeon"
	"github.com/cosmonaut/pigeon/x/pigeon/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PigeonKeeper(t)
	pigeon.InitGenesis(ctx, *k, genesisState)
	got := pigeon.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
