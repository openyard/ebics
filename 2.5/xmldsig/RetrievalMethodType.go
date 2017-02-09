// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type RetrievalMethodType struct {
	URI        w3c.AnyURI `xml:"URI,attr"`
	Type       w3c.AnyURI `xml:"Type,attr,omitempty"`
	Transforms Transforms `xml:"Transforms,omitempty"`
}
