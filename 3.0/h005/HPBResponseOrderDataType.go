// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h005

import esig "github.com/openyard/ebics/3.0/s002"
import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type HPBResponseOrderDataType struct {
	AuthenticationPubKeyInfo AuthenticationPubKeyInfoType `xml:"AuthenticationPubKeyInfo"`
	EncryptionPubKeyInfo     EncryptionPubKeyInfoType     `xml:"EncryptionPubKeyInfo"`
	HostID                   HostIDType                   `xml:"HostID"`
	SignaturePubKeyInfo      esig.SignaturePubKeyInfo     `xml:"http://www.ebics.org/S002 SignaturePubKeyInfo,omitempty"`

	Any []*w3c.Any
}
