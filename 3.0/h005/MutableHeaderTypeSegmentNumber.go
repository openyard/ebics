// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex element
type MutableHeaderTypeSegmentNumber struct {
	Value       *SegmentNumberType `xml:",chardata"`
	LastSegment *w3c.Boolean       `xml:"LastSegment,attr"`
}
