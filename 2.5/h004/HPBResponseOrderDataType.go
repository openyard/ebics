// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
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
