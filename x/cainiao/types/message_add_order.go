package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddOrder = "add_order"

var _ sdk.Msg = &MsgAddOrder{}

func NewMsgAddOrder(creator string, sendName string, sendAddress string, sendTel string, receiveName string, receiveAddress string, receiveTel string) *MsgAddOrder {
	return &MsgAddOrder{
		Creator:        creator,
		SendName:       sendName,
		SendAddress:    sendAddress,
		SendTel:        sendTel,
		ReceiveName:    receiveName,
		ReceiveAddress: receiveAddress,
		ReceiveTel:     receiveTel,
	}
}

func (msg *MsgAddOrder) Route() string {
	return RouterKey
}

func (msg *MsgAddOrder) Type() string {
	return TypeMsgAddOrder
}

func (msg *MsgAddOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
