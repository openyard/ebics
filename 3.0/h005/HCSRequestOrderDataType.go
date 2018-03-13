// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h005

import esig "github.com/openyard/ebics/3.0/s002"
import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type HCSRequestOrderDataType struct {
	AuthenticationPubKeyInfo AuthenticationPubKeyInfoType `xml:"AuthenticationPubKeyInfo"`
	EncryptionPubKeyInfo     EncryptionPubKeyInfoType     `xml:"EncryptionPubKeyInfo"`
	PartnerID                PartnerIDType                `xml:"PartnerID"`
	UserID                   UserIDType                   `xml:"UserID"`
	SignaturePubKeyInfo      esig.SignaturePubKeyInfo     `xml:"http://www.ebics.org/S002 SignaturePubKeyInfo"`

	Any []*w3c.Any
}
