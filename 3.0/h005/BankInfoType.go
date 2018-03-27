// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type BankInfoType struct {
	HostID    *HostIDType `xml:"HostID"`
	Parameter *Parameter  `xml:"Parameter,omitempty"`

	Any []*w3c.Any
}

func NewBankInfoType() *BankInfoType {
	return new(BankInfoType)
}

func (me *BankInfoType) SetHostID(value *HostIDType) {
	me.HostID = value
}

func (me *BankInfoType) AddHostID() *HostIDType {
	me.HostID = new(HostIDType)
	return me.HostID
}

func (me *BankInfoType) SetParameter(value *Parameter) {
	me.Parameter = value
}
