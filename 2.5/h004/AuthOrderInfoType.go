// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
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
