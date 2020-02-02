package types

import (
	"encoding/hex"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

// ensure Msg interface compliance at compile time
var (
	_ sdk.Msg = &MsgDelegateFeedConsent{}
	_ sdk.Msg = &MsgExchangeRatePrevote{}
	_ sdk.Msg = &MsgExchangeRateVote{}
)

//-------------------------------------------------
//-------------------------------------------------

// MsgExchangeRatePrevote - struct for prevoting on the ExchangeRateVote.
// The purpose of prevote is to hide vote exchange rate with hash
// which is formatted as hex string in SHA256("salt:exchange_rate:denom:voter")
type MsgExchangeRatePrevote struct {
	Hash      string         `json:"hash" yaml:"hash"` // hex string
	Denom     string         `json:"denom" yaml:"denom"`
	Feeder    sdk.AccAddress `json:"feeder" yaml:"feeder"`
	Validator sdk.ValAddress `json:"validator" yaml:"validator"`
}

// NewMsgExchangeRatePrevote creates a MsgExchangeRatePrevote instance
func NewMsgExchangeRatePrevote(VoteHash string, denom string, feederAddress sdk.AccAddress, valAddress sdk.ValAddress) MsgExchangeRatePrevote {
	return MsgExchangeRatePrevote{
		Hash:      VoteHash,
		Denom:     denom,
		Feeder:    feederAddress,
		Validator: valAddress,
	}
}

// Route implements sdk.Msg
func (msg MsgExchangeRatePrevote) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgExchangeRatePrevote) Type() string { return "exchangerateprevote" }

// GetSignBytes implements sdk.Msg
func (msg MsgExchangeRatePrevote) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements sdk.Msg
func (msg MsgExchangeRatePrevote) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Feeder}
}

// ValidateBasic Implements sdk.Msg
func (msg MsgExchangeRatePrevote) ValidateBasic() error {

	bz, err := hex.DecodeString(msg.Hash)
	if err != nil {
		return sdkerrors.Wrap(ErrInvalidHash, err.Error())
	}

	if l := len(bz); l != tmhash.TruncatedSize {
		return sdkerrors.Wrap(ErrInvalidHashLength, strconv.FormatInt(int64(l), 10))
	}

	if len(msg.Denom) == 0 {
		return ErrUnknowDenom
	}

	if msg.Feeder.Empty() {
		return sdkerrors.ErrInvalidAddress
	}

	if msg.Validator.Empty() {
		return sdkerrors.ErrInvalidAddress
	}

	return nil
}

// String implements fmt.Stringer interface
func (msg MsgExchangeRatePrevote) String() string {
	return fmt.Sprintf(`MsgExchangeRateVote
	hash:         %s,
	feeder:       %s, 
	validator:    %s, 
	denom:        %s`,
		msg.Hash, msg.Feeder, msg.Validator, msg.Denom)
}

// MsgExchangeRateVote - struct for voting on the exchange rate of Luna denominated in various Terra assets.
// For example, if the validator believes that the effective exchange rate of Luna in USD is 10.39, that's
// what the exchange rate field would be, and if 1213.34 for KRW, same.
type MsgExchangeRateVote struct {
	ExchangeRate sdk.Dec        `json:"exchange_rate" yaml:"exchange_rate"` // the effective rate of Luna in {Denom}
	Salt         string         `json:"salt" yaml:"salt"`
	Denom        string         `json:"denom" yaml:"denom"`
	Feeder       sdk.AccAddress `json:"feeder" yaml:"feeder"`
	Validator    sdk.ValAddress `json:"validator" yaml:"validator"`
}

// NewMsgExchangeRateVote creates a MsgExchangeRateVote instance
func NewMsgExchangeRateVote(rate sdk.Dec, salt string, denom string, feederAddress sdk.AccAddress, valAddress sdk.ValAddress) MsgExchangeRateVote {
	return MsgExchangeRateVote{
		ExchangeRate: rate,
		Salt:         salt,
		Denom:        denom,
		Feeder:       feederAddress,
		Validator:    valAddress,
	}
}

// Route implements sdk.Msg
func (msg MsgExchangeRateVote) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgExchangeRateVote) Type() string { return "exchangeratevote" }

// GetSignBytes implements sdk.Msg
func (msg MsgExchangeRateVote) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements sdk.Msg
func (msg MsgExchangeRateVote) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Feeder}
}

// ValidateBasic implements sdk.Msg
func (msg MsgExchangeRateVote) ValidateBasic() error {

	if len(msg.Denom) == 0 {
		return ErrUnknowDenom
	}

	if msg.Feeder.Empty() {
		return sdkerrors.ErrInvalidAddress
	}

	if msg.Validator.Empty() {
		return sdkerrors.ErrInvalidAddress
	}

	// Check overflow bit length
	if msg.ExchangeRate.BitLen() > 100+sdk.DecimalPrecisionBits {
		return sdkerrors.Wrap(ErrInvalidExchangeRate, msg.ExchangeRate.String())
	}

	if l := len(msg.Salt); l > 4 || l < 1 {
		return sdkerrors.Wrap(ErrInvalidSaltLength, strconv.FormatInt(int64(l), 10))
	}

	return nil
}

// String implements fmt.Stringer interface
func (msg MsgExchangeRateVote) String() string {
	return fmt.Sprintf(`MsgExchangeRateVote
	exchangerate:      %s,
	salt:       %s,
	feeder:     %s, 
	validator:  %s, 
	denom:      %s`,
		msg.ExchangeRate, msg.Salt, msg.Feeder, msg.Validator, msg.Denom)
}

// MsgDelegateFeedConsent - struct for delegating oracle voting rights to another address.
type MsgDelegateFeedConsent struct {
	Operator sdk.ValAddress `json:"operator" yaml:"operator"`
	Delegate sdk.AccAddress `json:"delegate" yaml:"delegate"`
}

// NewMsgDelegateFeedConsent creates a MsgDelegateFeedConsent instance
func NewMsgDelegateFeedConsent(operatorAddress sdk.ValAddress, feederAddress sdk.AccAddress) MsgDelegateFeedConsent {
	return MsgDelegateFeedConsent{
		Operator: operatorAddress,
		Delegate: feederAddress,
	}
}

// Route implements sdk.Msg
func (msg MsgDelegateFeedConsent) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgDelegateFeedConsent) Type() string { return "delegatefeeder" }

// GetSignBytes implements sdk.Msg
func (msg MsgDelegateFeedConsent) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements sdk.Msg
func (msg MsgDelegateFeedConsent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Operator)}
}

// ValidateBasic implements sdk.Msg
func (msg MsgDelegateFeedConsent) ValidateBasic() error {
	if msg.Operator.Empty() {
		return sdkerrors.ErrInvalidAddress
	}

	if msg.Delegate.Empty() {
		return sdkerrors.ErrInvalidAddress
	}

	return nil
}

// String implements fmt.Stringer interface
func (msg MsgDelegateFeedConsent) String() string {
	return fmt.Sprintf(`MsgDelegateFeedConsent
	operator:    %s, 
	delegate:   %s`,
		msg.Operator, msg.Delegate)
}
