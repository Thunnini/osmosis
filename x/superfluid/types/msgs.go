package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// constants
const (
	TypeMsgSuperfluidDelegate   = "superfluid_delegate"
	TypeMsgSuperfluidUndelegate = "superfluid_undelegate"
	TypeMsgSuperfluidRedelegate = "superfluid_redelegate"
)

var _ sdk.Msg = &MsgSuperfluidDelegate{}

// NewMsgSuperfluidDelegate creates a message to do superfluid delegation
func NewMsgSuperfluidDelegate(sender sdk.AccAddress, lockId uint64, valAddr sdk.ValAddress) *MsgSuperfluidDelegate {
	return &MsgSuperfluidDelegate{
		Sender:  sender.String(),
		LockId:  lockId,
		ValAddr: valAddr.String(),
	}
}

func (m MsgSuperfluidDelegate) Route() string { return RouterKey }
func (m MsgSuperfluidDelegate) Type() string  { return TypeMsgSuperfluidDelegate }
func (m MsgSuperfluidDelegate) ValidateBasic() error {
	if m.Sender == "" {
		return fmt.Errorf("sender should not be an empty address")
	}
	if m.LockId == 0 {
		return fmt.Errorf("lock id should be positive: %d < 0", m.LockId)
	}
	if m.ValAddr == "" {
		return fmt.Errorf("ValAddr should not be empty")
	}
	return nil
}
func (m MsgSuperfluidDelegate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}
func (m MsgSuperfluidDelegate) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgSuperfluidUndelegate{}

// NewMsgSuperfluidUndelegate creates a message to do superfluid undelegation
func NewMsgSuperfluidUndelegate(sender sdk.AccAddress, lockId uint64) *MsgSuperfluidUndelegate {
	return &MsgSuperfluidUndelegate{
		Sender: sender.String(),
		LockId: lockId,
	}
}

func (m MsgSuperfluidUndelegate) Route() string { return RouterKey }
func (m MsgSuperfluidUndelegate) Type() string  { return TypeMsgSuperfluidUndelegate }
func (m MsgSuperfluidUndelegate) ValidateBasic() error {
	if m.Sender == "" {
		return fmt.Errorf("sender should not be an empty address")
	}
	if m.LockId == 0 {
		return fmt.Errorf("lock id should be positive: %d < 0", m.LockId)
	}
	return nil
}
func (m MsgSuperfluidUndelegate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}
func (m MsgSuperfluidUndelegate) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgSuperfluidRedelegate{}

// NewMsgSuperfluidRedelegate creates a message to do superfluid redelegation
func NewMsgSuperfluidRedelegate(sender sdk.AccAddress, lockId uint64, newValAddr sdk.ValAddress) *MsgSuperfluidRedelegate {
	return &MsgSuperfluidRedelegate{
		Sender:     sender.String(),
		LockId:     lockId,
		NewValAddr: newValAddr.String(),
	}
}

func (m MsgSuperfluidRedelegate) Route() string { return RouterKey }
func (m MsgSuperfluidRedelegate) Type() string  { return TypeMsgSuperfluidRedelegate }
func (m MsgSuperfluidRedelegate) ValidateBasic() error {
	if m.Sender == "" {
		return fmt.Errorf("sender should not be an empty address")
	}
	if m.LockId == 0 {
		return fmt.Errorf("lock id should be positive: %d < 0", m.LockId)
	}
	if m.NewValAddr == "" {
		return fmt.Errorf("NewValAddr should not be empty")
	}
	return nil
}
func (m MsgSuperfluidRedelegate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}
func (m MsgSuperfluidRedelegate) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{sender}
}
