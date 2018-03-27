// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
