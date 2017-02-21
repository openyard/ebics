// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type HVUOrderDetailsType struct {
	OrderType      OrderTBaseType        `xml:"OrderType"`
	FileFormat     FileFormatType        `xml:"FileFormat,omitempty"`
	OrderID        OrderIDType           `xml:"OrderID"`
	OrderDataSize  w3c.PositiveInteger   `xml:"OrderDataSize"`
	SigningInfo    HVUSigningInfoType    `xml:"SigningInfo"`
	SignerInfo     SignerInfoType        `xml:"SignerInfo,omitempty"`
	OriginatorInfo HVUOriginatorInfoType `xml:"OriginatorInfo"`

	Any []*w3c.Any
}
