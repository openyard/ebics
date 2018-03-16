// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexElement
type AttributedAccountTypeAccountHolder struct {
	Value       *AccountHolderType     `xml:",chardata"`
	Role        *AccountHolderRoleType `xml:"Role,attr"`
	Description *w3c.NormalizedString  `xml:"Description,attr"`
}

func NewAttributedAccountTypeAccountHolder() *AttributedAccountTypeAccountHolder {
	return new(AttributedAccountTypeAccountHolder)
}

func (me *AttributedAccountTypeAccountHolder) SetRole(value *AccountHolderRoleType) {
	me.Role = value
}

func (me *AttributedAccountTypeAccountHolder) SetDescription(value *w3c.NormalizedString) {
	me.Description = value
}
