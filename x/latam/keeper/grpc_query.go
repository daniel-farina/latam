package keeper

import (
	"github.com/daniel-farina/latam/x/latam/types"
)

var _ types.QueryServer = Keeper{}
