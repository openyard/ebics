// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type SignerInfoType struct {
	PartnerID  PartnerIDType            `xml:"PartnerID"`
	UserID     UserIDType               `xml:"UserID"`
	Name       NameType                 `xml:"Name,omitempty"`
	Timestamp  TimestampType            `xml:"Timestamp"`
	Permission SignerInfoTypePermission `xml:"Permission"`

	Any []*w3c.Any
}
