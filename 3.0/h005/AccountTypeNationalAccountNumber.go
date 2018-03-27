// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexElement
type AccountTypeNationalAccountNumber struct {
	Value  *NationalAccountNumberType `xml:",chardata"`
	Format *w3c.Token                 `xml:"Format,attr"`
}

func NewAccountTypeNationalAccountNumber() *AccountTypeNationalAccountNumber {
	return new(AccountTypeNationalAccountNumber)
}

func (me *AccountTypeNationalAccountNumber) SetFormat(value *w3c.Token) {
	me.Format = value
}
