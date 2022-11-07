package globalfee

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/CosmosContracts/juno/v11/x/globalfee/types"
)

var _ types.QueryServer = &Querier{}

type Querier struct {
	paramSource paramSource
}

func NewQuerier(paramSource paramSource) Querier {
	return Querier{paramSource: paramSource}
}

// MinimumGasPrices return minimum gas prices
func (g Querier) MinimumGasPrices(stdCtx context.Context, _ *types.QueryMinimumGasPricesRequest) (*types.QueryMinimumGasPricesResponse, error) {
	var minGasPrices sdk.DecCoins
	ctx := sdk.UnwrapSDKContext(stdCtx)
	if g.paramSource.Has(ctx, types.ParamStoreKeyMinGasPrices) {
		g.paramSource.Get(ctx, types.ParamStoreKeyMinGasPrices, &minGasPrices)
	}
	return &types.QueryMinimumGasPricesResponse{
		MinimumGasPrices: minGasPrices,
	}, nil
}

// WhiteList return wallet accounts allowed to override fees
func (g Querier) WhiteList(c context.Context, _ *types.QueryWhiteListRequest) (*types.QueryWhiteListResponse, error) {
	var accounts types.AccountRecords
	ctx := sdk.UnwrapSDKContext(c)
	if g.paramSource.Has(ctx, types.ParamStoreKeyMinGasPrices) {
		g.paramSource.Get(ctx, types.ParamStoreKeyMinGasPrices, &accounts)
	}
	return &types.QueryWhiteListResponse{AccountRecord: accounts}, nil
}
