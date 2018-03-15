// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type HVUOrderDetailsType struct {
	Service             *RestrictedServiceType `xml:"Service"`
	OrderID             *OrderIDType           `xml:"OrderID"`
	OrderDataSize       *w3c.PositiveInteger   `xml:"OrderDataSize"`
	SigningInfo         *HVUSigningInfoType    `xml:"SigningInfo"`
	SignerInfo          []*SignerInfoType      `xml:"SignerInfo,omitempty"`
	OriginatorInfo      *HVUOriginatorInfoType `xml:"OriginatorInfo"`
	AdditionalOrderInfo *String255Type         `xml:"AdditionalOrderInfo,omitempty"`

	Any []*w3c.Any
}
