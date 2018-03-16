// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

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
