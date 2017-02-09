// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type HVUOriginatorInfoType struct {
	PartnerID PartnerIDType `xml:"PartnerID"`
	UserID    UserIDType    `xml:"UserID"`
	Name      NameType      `xml:"Name,omitempty"`
	Timestamp TimestampType `xml:"Timestamp"`

	Any []*w3c.Any
}
