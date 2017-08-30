// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type AddressInfoType struct {
	Name     NameType  `xml:"Name,omitempty"`
	Street   NameType  `xml:"Street,omitempty"`
	PostCode w3c.Token `xml:"PostCode,omitempty"`
	City     NameType  `xml:"City,omitempty"`
	Region   NameType  `xml:"Region,omitempty"`
	Country  NameType  `xml:"Country,omitempty"`

	Any []*w3c.Any
}
