// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
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
