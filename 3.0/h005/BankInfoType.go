// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *BankInfoType) SetParameter(value *Parameter) {
	me.Parameter = value
}
