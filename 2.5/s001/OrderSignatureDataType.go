// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
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
