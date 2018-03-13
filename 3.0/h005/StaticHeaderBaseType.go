// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type StaticHeaderBaseType struct {
	HostID         HostIDType         `xml:"HostID"`
	Nonce          NonceType          `xml:"Nonce,omitempty"`
	Timestamp      TimestampType      `xml:"Timestamp,omitempty"`
	PartnerID      PartnerIDType      `xml:"PartnerID"`
	UserID         UserIDType         `xml:"UserID"`
	SystemID       UserIDType         `xml:"SystemID,omitempty"`
	Product        ProductElementType `xml:"Product,omitempty"`
	OrderDetails   OrderDetailsType   `xml:"OrderDetails"`
	SecurityMedium SecurityMediumType `xml:"SecurityMedium"`

	Any []*w3c.Any
}
