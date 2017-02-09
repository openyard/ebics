// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

import esig "github.com/openyard/ebics/2.5/s001"
import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type HPBResponseOrderDataType struct {
	AuthenticationPubKeyInfo AuthenticationPubKeyInfoType `xml:"AuthenticationPubKeyInfo"`
	EncryptionPubKeyInfo     EncryptionPubKeyInfoType     `xml:"EncryptionPubKeyInfo"`
	HostID                   HostIDType                   `xml:"HostID"`
	SignaturePubKeyInfo      esig.SignaturePubKeyInfo     `xml:"http://www.ebics.org/S001 SignaturePubKeyInfo,omitempty"`

	Any []*w3c.Any
}
