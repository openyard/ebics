// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
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
