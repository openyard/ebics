// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexElement
type AccountTypeBankCode struct {
	Value         *BankCodeType       `xml:",chardata"`
	International *w3c.Boolean        `xml:"International,attr,omitempty"`
	Prefix        *BankCodePrefixType `xml:"Prefix,attr,omitempty"`
}

func NewAccountTypeBankCode() *AccountTypeBankCode {
	return new(AccountTypeBankCode)
}

func (me *AccountTypeBankCode) SetInternational(value *w3c.Boolean) {
	me.International = value
}

func (me *AccountTypeBankCode) SetPrefix(value *BankCodePrefixType) {
	me.Prefix = value
}
