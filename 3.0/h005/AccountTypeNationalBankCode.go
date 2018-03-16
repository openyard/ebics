// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexElement
type AccountTypeNationalBankCode struct {
	Value  *NationalBankCodeType `xml:",chardata"`
	Format *w3c.Token            `xml:"Format,attr"`
}

func NewAccountTypeNationalBankCode() *AccountTypeNationalBankCode {
	return new(AccountTypeNationalBankCode)
}

func (me *AccountTypeNationalBankCode) SetFormat(value *w3c.Token) {
	me.Format = value
}
