// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
