// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type ObjectType struct {
	Id       w3c.ID     `xml:"Id,attr,omitempty"`
	MimeType w3c.String `xml:"MimeType,attr,omitempty"`
	Encoding w3c.AnyURI `xml:"Encoding,attr,omitempty"`

	Any []*w3c.Any
}
