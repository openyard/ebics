// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexElement
type AccountTypeAccountNumber struct {
	Value         *AccountNumberType `xml:",chardata"`
	International *w3c.Boolean       `xml:"International,attr,omitempty"`
}

func NewAccountTypeAccountNumber() *AccountTypeAccountNumber {
	return new(AccountTypeAccountNumber)
}

func (me *AccountTypeAccountNumber) SetInternational(value *w3c.Boolean) {
	me.International = value
}
