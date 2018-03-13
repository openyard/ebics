// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type ReferenceType struct {
	Id           w3c.ID       `xml:"Id,attr,omitempty"`
	URI          w3c.AnyURI   `xml:"URI,attr,omitempty"`
	Type         w3c.AnyURI   `xml:"Type,attr,omitempty"`
	Transforms   Transforms   `xml:"Transforms,omitempty"`
	DigestMethod DigestMethod `xml:"DigestMethod"`
	DigestValue  DigestValue  `xml:"DigestValue"`
}
