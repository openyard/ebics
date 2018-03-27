// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexElement
type AttributedAccountTypeBankCode struct {
	Value         *BankCodeType         `xml:",chardata"`
	Role          *BankCodeRoleType     `xml:"Role,attr"`
	Description   *w3c.NormalizedString `xml:"Description,attr"`
	International *w3c.Boolean          `xml:"International,attr,omitempty"`
	Prefix        *BankCodePrefixType   `xml:"Prefix,attr,omitempty"`
}

func NewAttributedAccountTypeBankCode() *AttributedAccountTypeBankCode {
	return new(AttributedAccountTypeBankCode)
}

func (me *AttributedAccountTypeBankCode) SetRole(value *BankCodeRoleType) {
	me.Role = value
}

func (me *AttributedAccountTypeBankCode) SetDescription(value *w3c.NormalizedString) {
	me.Description = value
}

func (me *AttributedAccountTypeBankCode) SetInternational(value *w3c.Boolean) {
	me.International = value
}

func (me *AttributedAccountTypeBankCode) SetPrefix(value *BankCodePrefixType) {
	me.Prefix = value
}
