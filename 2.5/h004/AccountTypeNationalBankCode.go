// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

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
