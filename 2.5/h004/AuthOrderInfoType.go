// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type AuthOrderInfoType struct {
	OrderType      OrderTBaseType         `xml:"OrderType"`
	FileFormat     FileFormatType         `xml:"FileFormat,omitempty"`
	TransferType   TransferType           `xml:"TransferType"`
	OrderFormat    OrderFormatType        `xml:"OrderFormat,omitempty"`
	Description    OrderDescriptionType   `xml:"Description"`
	NumSigRequired w3c.NonNegativeInteger `xml:"NumSigRequired,omitempty"`

	Any []*w3c.Any
}
