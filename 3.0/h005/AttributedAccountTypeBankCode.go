// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexElement
type AttributedAccountTypeBankCode struct {
	Value         *BankCodeType         `xml:",chardata"`
	Role          *BankCodeRoleType     `xml:"Role,attr"`
	Description   *w3c.NormalizedString `xml:"Description,attr"`
	International *w3c.Boolean          `xml:"International,attr,omitempty"`
	Prefix        *BankCodePrefixType   `xml:"Prefix,attr,omitempty"`
}
