// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type SignerInfoType struct {
	PartnerID  *PartnerIDType            `xml:"PartnerID"`
	UserID     *UserIDType               `xml:"UserID"`
	Name       *NameType                 `xml:"Name,omitempty"`
	Timestamp  *TimestampType            `xml:"Timestamp"`
	Permission *SignerInfoTypePermission `xml:"Permission"`

	Any []*w3c.Any
}
