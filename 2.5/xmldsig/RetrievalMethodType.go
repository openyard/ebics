// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type RetrievalMethodType struct {
	URI        w3c.AnyURI `xml:"URI,attr"`
	Type       w3c.AnyURI `xml:"Type,attr,omitempty"`
	Transforms Transforms `xml:"Transforms,omitempty"`
}
