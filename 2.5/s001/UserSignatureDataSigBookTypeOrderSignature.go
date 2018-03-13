// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package s001

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex element
type UserSignatureDataSigBookTypeOrderSignature struct {
	Value     OrderSignatureType `xml:",chardata"`
	PartnerID w3c.String         `xml:"PartnerID,attr"`
}
