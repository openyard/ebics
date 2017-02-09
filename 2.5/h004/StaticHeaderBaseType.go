// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

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
