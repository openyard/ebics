// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

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
