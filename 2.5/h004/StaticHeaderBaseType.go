// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type StaticHeaderBaseType struct {
	HostID         *HostIDType         `xml:"HostID"`
	Nonce          *NonceType          `xml:"Nonce,omitempty"`
	Timestamp      *TimestampType      `xml:"Timestamp,omitempty"`
	PartnerID      *PartnerIDType      `xml:"PartnerID"`
	UserID         *UserIDType         `xml:"UserID"`
	SystemID       *UserIDType         `xml:"SystemID,omitempty"`
	Product        *ProductElementType `xml:"Product,omitempty"`
	OrderDetails   *OrderDetailsType   `xml:"OrderDetails"`
	SecurityMedium *SecurityMediumType `xml:"SecurityMedium"`

	Any []*w3c.Any
}
