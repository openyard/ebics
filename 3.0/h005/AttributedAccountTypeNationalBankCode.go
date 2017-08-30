// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex element
type AttributedAccountTypeNationalBankCode struct {
	Value       NationalBankCodeType `xml:",chardata"`
	Role        BankCodeRoleType     `xml:"Role,attr"`
	Description w3c.NormalizedString `xml:"Description,attr"`
	Format      w3c.Token            `xml:"Format,attr"`
}
