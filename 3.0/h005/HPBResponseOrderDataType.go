// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
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
