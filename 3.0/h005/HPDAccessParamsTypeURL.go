// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex element
type HPDAccessParamsTypeURL struct {
	Value      w3c.AnyURI    `xml:",chardata"`
	Valid_from TimestampType `xml:"Valid_from,attr"`
}
