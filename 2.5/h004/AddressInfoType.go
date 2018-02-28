// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

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
