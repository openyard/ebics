// Generated with goxc v - rev
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex element
type ParameterValue struct {
	Value w3c.AnySimpleType `xml:",chardata"`
	Type  w3c.NCName        `xml:"Type,attr"`
}
