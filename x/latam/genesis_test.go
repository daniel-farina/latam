package latam_test

import (
	"testing"

	keepertest "github.com/daniel-farina/latam/testutil/keeper"
	"github.com/daniel-farina/latam/x/latam"
	"github.com/daniel-farina/latam/x/latam/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LatamKeeper(t)
	latam.InitGenesis(ctx, *k, genesisState)
	got := latam.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
