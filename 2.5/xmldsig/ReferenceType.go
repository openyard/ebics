// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
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
