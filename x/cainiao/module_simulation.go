package cainiao

import (
	"math/rand"

	"github.com/cosmonaut/cainiao/testutil/sample"
	cainiaosimulation "github.com/cosmonaut/cainiao/x/cainiao/simulation"
	"github.com/cosmonaut/cainiao/x/cainiao/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = cainiaosimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgAddOrder = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddOrder int = 100

	opWeightMsgUpdateOrder = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateOrder int = 100

	opWeightMsgUpdateOrderState = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateOrderState int = 100

	opWeightMsgReceiveOrder = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgReceiveOrder int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	cainiaoGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cainiaoGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgAddOrder int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddOrder, &weightMsgAddOrder, nil,
		func(_ *rand.Rand) {
			weightMsgAddOrder = defaultWeightMsgAddOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddOrder,
		cainiaosimulation.SimulateMsgAddOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateOrder int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateOrder, &weightMsgUpdateOrder, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateOrder = defaultWeightMsgUpdateOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateOrder,
		cainiaosimulation.SimulateMsgUpdateOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateOrderState int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateOrderState, &weightMsgUpdateOrderState, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateOrderState = defaultWeightMsgUpdateOrderState
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateOrderState,
		cainiaosimulation.SimulateMsgUpdateOrderState(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgReceiveOrder int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgReceiveOrder, &weightMsgReceiveOrder, nil,
		func(_ *rand.Rand) {
			weightMsgReceiveOrder = defaultWeightMsgReceiveOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgReceiveOrder,
		cainiaosimulation.SimulateMsgReceiveOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
