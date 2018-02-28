// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package h005

import ds "github.com/openyard/ebics/3.0/xmldsig"

// complex type
type EbicsNoPubKeyDigestsRequestBody struct {
	X509Data ds.X509Data `xml:"http://www.w3.org/2000/09/xmldsig# X509Data,omitempty"`
}
