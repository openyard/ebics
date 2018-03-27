// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

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
