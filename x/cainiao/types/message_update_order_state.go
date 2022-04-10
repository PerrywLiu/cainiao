package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateOrderState = "update_order_state"

var _ sdk.Msg = &MsgUpdateOrderState{}

func NewMsgUpdateOrderState(creator string, id string, station string) *MsgUpdateOrderState {
	return &MsgUpdateOrderState{
		Creator: creator,
		Id:      id,
		Station: station,
	}
}

func (msg *MsgUpdateOrderState) Route() string {
	return RouterKey
}

func (msg *MsgUpdateOrderState) Type() string {
	return TypeMsgUpdateOrderState
}

func (msg *MsgUpdateOrderState) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateOrderState) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateOrderState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
