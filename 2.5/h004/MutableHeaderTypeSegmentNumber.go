// Generated with goxc v - rev
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex element
type MutableHeaderTypeSegmentNumber struct {
	Value       SegmentNumberType `xml:",chardata"`
	LastSegment w3c.Boolean       `xml:"LastSegment,attr"`
}
