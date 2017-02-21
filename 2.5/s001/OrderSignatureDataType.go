// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
package s001

import w3c "github.com/openyard/ebics/2.5/w3c"
import ds "github.com/openyard/ebics/2.5/xmldsig"

// complex type
type OrderSignatureDataType struct {
	SignatureVersion SignatureVersionType `xml:"SignatureVersion"`
	SignatureValue   w3c.Base64Binary     `xml:"SignatureValue"`
	PartnerID        PartnerIDType        `xml:"PartnerID"`
	UserID           UserIDType           `xml:"UserID"`
	X509Data         ds.X509Data          `xml:"http://www.w3.org/2000/09/xmldsig# X509Data,omitempty"`
}
