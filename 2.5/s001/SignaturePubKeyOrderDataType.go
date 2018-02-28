// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package s001

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type SignaturePubKeyOrderDataType struct {
	SignaturePubKeyInfo SignaturePubKeyInfoType `xml:"SignaturePubKeyInfo"`
	PartnerID           PartnerIDType           `xml:"PartnerID"`
	UserID              UserIDType              `xml:"UserID"`

	Any []*w3c.Any
}
