// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
package h004

import ds "github.com/openyard/ebics/2.5/xmldsig"

// complex type
type EbicsNoPubKeyDigestsRequestBody struct {
	X509Data ds.X509Data `xml:"http://www.w3.org/2000/09/xmldsig# X509Data,omitempty"`
}
