// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type SignatureMethodType struct {
	Algorithm        w3c.AnyURI           `xml:"Algorithm,attr"`
	HMACOutputLength HMACOutputLengthType `xml:"HMACOutputLength,omitempty"`

	Any []*w3c.Any
}
