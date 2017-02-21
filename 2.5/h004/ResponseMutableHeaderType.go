// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type ResponseMutableHeaderType struct {
	TransactionPhase TransactionPhaseType                   `xml:"TransactionPhase"`
	SegmentNumber    ResponseMutableHeaderTypeSegmentNumber `xml:"SegmentNumber,omitempty"`
	OrderID          OrderIDType                            `xml:"OrderID,omitempty"`
	ReturnCode       ReturnCodeType                         `xml:"ReturnCode"`
	ReportText       ReportTextType                         `xml:"ReportText"`

	Any []*w3c.Any
}
