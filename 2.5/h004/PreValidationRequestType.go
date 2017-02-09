// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type PreValidationRequestType struct {
	DataDigest           DataDigestType               `xml:"DataDigest,omitempty"`
	AccountAuthorisation PreValidationAccountAuthType `xml:"AccountAuthorisation,omitempty"`

	Any []*w3c.Any
}
