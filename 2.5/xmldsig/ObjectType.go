// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type ObjectType struct {
	Id       w3c.ID     `xml:"Id,attr,omitempty"`
	MimeType w3c.String `xml:"MimeType,attr,omitempty"`
	Encoding w3c.AnyURI `xml:"Encoding,attr,omitempty"`

	Any []*w3c.Any
}
