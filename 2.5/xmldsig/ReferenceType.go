// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type ReferenceType struct {
	Id           *w3c.ID       `xml:"Id,attr,omitempty"`
	URI          *w3c.AnyURI   `xml:"URI,attr,omitempty"`
	Type         *w3c.AnyURI   `xml:"Type,attr,omitempty"`
	Transforms   *Transforms   `xml:"Transforms,omitempty"`
	DigestMethod *DigestMethod `xml:"DigestMethod"`
	DigestValue  *DigestValue  `xml:"DigestValue"`
}
