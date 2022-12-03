package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	authforktypes "github.com/CosmosContracts/juno/v12/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// FeeSharePayoutDecorator Run his after we already deduct the fee from the account with
// the ante.NewDeductFeeDecorator() decorator. We pull funds from the FeeCollector ModuleAccount
type FeeSharePayoutDecorator struct {
	ak             authforktypes.AccountKeeper
	bankKeeper     authforktypes.BankKeeper
	feegrantKeeper authforktypes.FeegrantKeeper
	feesharekeeper authforktypes.FeeShareKeeper
}

func NewFeeSharePayoutDecorator(ak authforktypes.AccountKeeper, bk authforktypes.BankKeeper, fk authforktypes.FeegrantKeeper, fs authforktypes.FeeShareKeeper) FeeSharePayoutDecorator {
	return FeeSharePayoutDecorator{
		ak:             ak,
		bankKeeper:     bk,
		feegrantKeeper: fk,
		feesharekeeper: fs,
	}
}

func (fsd FeeSharePayoutDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	err = FeeSharePayout(ctx, fsd.bankKeeper, feeTx.GetFee(), fsd.feesharekeeper, tx.GetMsgs())
	if err != nil {
		return ctx, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, err.Error())
	}

	return next(ctx, tx, simulate)
}

type PayoutContracts struct {
	contractAddr sdk.AccAddress
	withdrawAddr sdk.AccAddress
}

// FeeSharePayout takes the total fees and redistributes 50% (or param set) to the contract developers
// provided they opted-in to payments.
func FeeSharePayout(ctx sdk.Context, bankKeeper authforktypes.BankKeeper, fees sdk.Coins, revKeeper authforktypes.FeeShareKeeper, msgs []sdk.Msg) error {
	// Pairs of contract address and withdraw address if one is set. If not, we don't place it in here
	pairs := make([]PayoutContracts, 0)
	for _, msg := range msgs {
		if _, ok := msg.(*wasmtypes.MsgExecuteContract); ok {
			contractAddr, err := sdk.AccAddressFromBech32(msg.(*wasmtypes.MsgExecuteContract).Contract)
			if err != nil {
				return err
			}

			shareData, _ := revKeeper.GetFeeShare(ctx, contractAddr)
			withdrawAddr := shareData.GetWithdrawerAddr()
			if withdrawAddr != nil {
				pairs = append(pairs, PayoutContracts{contractAddr, withdrawAddr})
			}
		}
	}

	// FeeShare logic payouts for contracts
	numPairs := len(pairs)
	if numPairs > 0 {
		govPercent := revKeeper.GetParams(ctx).DeveloperShares // 0.500000000

		// multiply times 100 = 50. This way we can do 100/50 = 2 for the split fee amount
		// if above is 25%, then 100/25 = 4 = they get 1/4th of the total fee between contracts
		splitNumber := govPercent.MulInt64(100)

		// Gets XX% of the fees based off governance params based off the number of contracts we execute on
		// (majority of the time this is only 1 contract). Should we simplify and only get the first contract?

		var splitFees sdk.Coins
		for _, c := range fees {
			divisor := sdk.NewDec(100).Quo(splitNumber).RoundInt64()
			// Ex: 10 coins / (100/50=2) = 5 coins is 50%
			// So each contract gets 5 coins / 2 contracts = 2.5 coins per. Does it round up?
			// This means if multiple contracts are in the message, the community pool may get less than 50% for these edge cases.
			reward := sdk.NewCoin(c.Denom, c.Amount.QuoRaw(divisor).QuoRaw(int64(numPairs)))
			if !reward.Amount.IsZero() {
				splitFees = append(splitFees, reward)
			}
		}

		// pay fees evenly between pairs of contracts
		for _, p := range pairs {
			err := bankKeeper.SendCoinsFromModuleToAccount(ctx, types.FeeCollectorName, p.withdrawAddr, splitFees)
			if err != nil {
				return sdkerrors.Wrapf(sdkerrors.Error{}, err.Error())
			}
		}
	}

	return nil
}