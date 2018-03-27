// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexElement
type AttributedAccountTypeNationalAccountNumber struct {
	Value       *NationalAccountNumberType `xml:",chardata"`
	Role        *AccountNumberRoleType     `xml:"Role,attr"`
	Description *w3c.NormalizedString      `xml:"Description,attr"`
	Format      *w3c.Token                 `xml:"Format,attr"`
}

func NewAttributedAccountTypeNationalAccountNumber() *AttributedAccountTypeNationalAccountNumber {
	return new(AttributedAccountTypeNationalAccountNumber)
}

func (me *AttributedAccountTypeNationalAccountNumber) SetRole(value *AccountNumberRoleType) {
	me.Role = value
}

func (me *AttributedAccountTypeNationalAccountNumber) SetDescription(value *w3c.NormalizedString) {
	me.Description = value
}

func (me *AttributedAccountTypeNationalAccountNumber) SetFormat(value *w3c.Token) {
	me.Format = value
}
