// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type ObjectType struct {
	Id       *w3c.ID     `xml:"Id,attr,omitempty"`
	MimeType *w3c.String `xml:"MimeType,attr,omitempty"`
	Encoding *w3c.AnyURI `xml:"Encoding,attr,omitempty"`

	Any []*w3c.Any
}
