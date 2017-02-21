// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
package s001

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex element
type UserSignatureDataSigBookTypeOrderSignature struct {
	Value     OrderSignatureType `xml:",chardata"`
	PartnerID w3c.String         `xml:"PartnerID,attr"`
}
