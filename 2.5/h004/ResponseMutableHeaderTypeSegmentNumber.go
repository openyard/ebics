// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexElement
type ResponseMutableHeaderTypeSegmentNumber struct {
	Value       *SegmentNumberType `xml:",chardata"`
	LastSegment *w3c.Boolean       `xml:"LastSegment,attr"`
}

func NewResponseMutableHeaderTypeSegmentNumber() *ResponseMutableHeaderTypeSegmentNumber {
	return new(ResponseMutableHeaderTypeSegmentNumber)
}

func (me *ResponseMutableHeaderTypeSegmentNumber) SetLastSegment(value *w3c.Boolean) {
	me.LastSegment = value
}
