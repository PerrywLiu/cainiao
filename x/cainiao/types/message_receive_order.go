package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgReceiveOrder = "receive_order"

var _ sdk.Msg = &MsgReceiveOrder{}

func NewMsgReceiveOrder(creator string, id int32) *MsgReceiveOrder {
	return &MsgReceiveOrder{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgReceiveOrder) Route() string {
	return RouterKey
}

func (msg *MsgReceiveOrder) Type() string {
	return TypeMsgReceiveOrder
}

func (msg *MsgReceiveOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgReceiveOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReceiveOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
