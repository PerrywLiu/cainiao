package simulation

import (
	"math/rand"

	"github.com/cosmonaut/cainiao/x/cainiao/keeper"
	"github.com/cosmonaut/cainiao/x/cainiao/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgAddOrder(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAddOrder{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AddOrder simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "AddOrder simulation not implemented"), nil, nil
	}
}
