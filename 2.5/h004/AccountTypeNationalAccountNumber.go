// Generated with goxc v - rev
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex element
type AccountTypeNationalAccountNumber struct {
	Value  NationalAccountNumberType `xml:",chardata"`
	Format w3c.Token                 `xml:"Format,attr"`
}
