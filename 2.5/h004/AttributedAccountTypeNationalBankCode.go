// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexElement
type AttributedAccountTypeNationalBankCode struct {
	Value       *NationalBankCodeType `xml:",chardata"`
	Role        *BankCodeRoleType     `xml:"Role,attr"`
	Description *w3c.NormalizedString `xml:"Description,attr"`
	Format      *w3c.Token            `xml:"Format,attr"`
}

func NewAttributedAccountTypeNationalBankCode() *AttributedAccountTypeNationalBankCode {
	return new(AttributedAccountTypeNationalBankCode)
}

func (me *AttributedAccountTypeNationalBankCode) SetRole(value *BankCodeRoleType) {
	me.Role = value
}

func (me *AttributedAccountTypeNationalBankCode) SetDescription(value *w3c.NormalizedString) {
	me.Description = value
}

func (me *AttributedAccountTypeNationalBankCode) SetFormat(value *w3c.Token) {
	me.Format = value
}
