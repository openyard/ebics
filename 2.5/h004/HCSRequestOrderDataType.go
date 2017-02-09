// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

import esig "github.com/openyard/ebics/2.5/s001"
import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type HCSRequestOrderDataType struct {
	AuthenticationPubKeyInfo AuthenticationPubKeyInfoType `xml:"AuthenticationPubKeyInfo"`
	EncryptionPubKeyInfo     EncryptionPubKeyInfoType     `xml:"EncryptionPubKeyInfo"`
	PartnerID                PartnerIDType                `xml:"PartnerID"`
	UserID                   UserIDType                   `xml:"UserID"`
	SignaturePubKeyInfo      esig.SignaturePubKeyInfo     `xml:"http://www.ebics.org/S001 SignaturePubKeyInfo"`

	Any []*w3c.Any
}
