// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type KeyMgmntResponseMutableHeaderType struct {
	OrderID    OrderIDType    `xml:"OrderID,omitempty"`
	ReturnCode ReturnCodeType `xml:"ReturnCode"`
	ReportText ReportTextType `xml:"ReportText"`

	Any []*w3c.Any
}
