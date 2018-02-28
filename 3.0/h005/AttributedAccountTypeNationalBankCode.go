// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex element
type AttributedAccountTypeNationalBankCode struct {
	Value       NationalBankCodeType `xml:",chardata"`
	Role        BankCodeRoleType     `xml:"Role,attr"`
	Description w3c.NormalizedString `xml:"Description,attr"`
	Format      w3c.Token            `xml:"Format,attr"`
}
