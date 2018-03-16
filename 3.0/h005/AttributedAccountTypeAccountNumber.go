// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexElement
type AttributedAccountTypeAccountNumber struct {
	Value         *AccountNumberType     `xml:",chardata"`
	Role          *AccountNumberRoleType `xml:"Role,attr"`
	Description   *w3c.NormalizedString  `xml:"Description,attr"`
	International *w3c.Boolean           `xml:"International,attr,omitempty"`
}

func NewAttributedAccountTypeAccountNumber() *AttributedAccountTypeAccountNumber {
	return new(AttributedAccountTypeAccountNumber)
}

func (me *AttributedAccountTypeAccountNumber) SetRole(value *AccountNumberRoleType) {
	me.Role = value
}

func (me *AttributedAccountTypeAccountNumber) SetDescription(value *w3c.NormalizedString) {
	me.Description = value
}

func (me *AttributedAccountTypeAccountNumber) SetInternational(value *w3c.Boolean) {
	me.International = value
}
