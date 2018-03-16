// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
