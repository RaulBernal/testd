package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/BitCannaGlobal/test/testutil/keeper"
	"github.com/BitCannaGlobal/test/x/test/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.TestKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
