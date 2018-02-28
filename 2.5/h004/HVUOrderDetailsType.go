// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
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
