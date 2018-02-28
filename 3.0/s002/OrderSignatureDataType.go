// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package s002

import w3c "github.com/openyard/ebics/3.0/w3c"
import ds "github.com/openyard/ebics/3.0/xmldsig"

// complex type
type OrderSignatureDataType struct {
	SignatureVersion SignatureVersionType `xml:"SignatureVersion"`
	SignatureValue   w3c.Base64Binary     `xml:"SignatureValue"`
	PartnerID        PartnerIDType        `xml:"PartnerID"`
	UserID           UserIDType           `xml:"UserID"`
	X509Data         ds.X509Data          `xml:"http://www.w3.org/2000/09/xmldsig# X509Data,omitempty"`
}
