// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type HPDVersionType struct {
	Protocol       string `xml:"Protocol"`
	Authentication string `xml:"Authentication"`
	Encryption     string `xml:"Encryption"`
	Signature      string `xml:"Signature"`

	Any []*w3c.Any
}
