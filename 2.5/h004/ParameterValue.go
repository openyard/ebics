// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex element
type ParameterValue struct {
	Value w3c.AnySimpleType `xml:",chardata"`
	Type  w3c.NCName        `xml:"Type,attr"`
}
