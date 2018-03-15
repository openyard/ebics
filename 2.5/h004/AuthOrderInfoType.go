// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type AuthOrderInfoType struct {
	OrderType      *OrderTBaseType         `xml:"OrderType"`
	FileFormat     *FileFormatType         `xml:"FileFormat,omitempty"`
	TransferType   *TransferType           `xml:"TransferType"`
	OrderFormat    *OrderFormatType        `xml:"OrderFormat,omitempty"`
	Description    *OrderDescriptionType   `xml:"Description"`
	NumSigRequired *w3c.NonNegativeInteger `xml:"NumSigRequired,omitempty"`

	Any []*w3c.Any
}
