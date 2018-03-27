// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

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
