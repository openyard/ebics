// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type HVUOrderDetailsType struct {
	OrderType      *OrderTBaseType        `xml:"OrderType"`
	FileFormat     *FileFormatType        `xml:"FileFormat,omitempty"`
	OrderID        *OrderIDType           `xml:"OrderID"`
	OrderDataSize  *w3c.PositiveInteger   `xml:"OrderDataSize"`
	SigningInfo    *HVUSigningInfoType    `xml:"SigningInfo"`
	SignerInfo     []*SignerInfoType      `xml:"SignerInfo,omitempty"`
	OriginatorInfo *HVUOriginatorInfoType `xml:"OriginatorInfo"`

	Any []*w3c.Any
}
