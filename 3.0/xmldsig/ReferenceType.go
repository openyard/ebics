// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type ReferenceType struct {
	Id           w3c.ID       `xml:"Id,attr,omitempty"`
	URI          w3c.AnyURI   `xml:"URI,attr,omitempty"`
	Type         w3c.AnyURI   `xml:"Type,attr,omitempty"`
	Transforms   Transforms   `xml:"Transforms,omitempty"`
	DigestMethod DigestMethod `xml:"DigestMethod"`
	DigestValue  DigestValue  `xml:"DigestValue"`
}
