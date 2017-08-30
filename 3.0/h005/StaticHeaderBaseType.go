// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
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
