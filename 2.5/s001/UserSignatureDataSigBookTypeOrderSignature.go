// Generated with goxc v - rev
package s001

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex element
type UserSignatureDataSigBookTypeOrderSignature struct {
	Value     OrderSignatureType `xml:",chardata"`
	PartnerID w3c.String         `xml:"PartnerID,attr"`
}
