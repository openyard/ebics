// Generated with goxc v - rev
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex element
type AttributedAccountTypeNationalBankCode struct {
	Value       NationalBankCodeType `xml:",chardata"`
	Role        BankCodeRoleType     `xml:"Role,attr"`
	Description w3c.NormalizedString `xml:"Description,attr"`
	Format      w3c.Token            `xml:"Format,attr"`
}
