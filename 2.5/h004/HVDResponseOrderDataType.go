// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type HVDResponseOrderDataType struct {
	DataDigest            DataDigestType      `xml:"DataDigest"`
	DisplayFile           w3c.Base64Binary    `xml:"DisplayFile"`
	OrderDataAvailable    w3c.Boolean         `xml:"OrderDataAvailable"`
	OrderDataSize         w3c.PositiveInteger `xml:"OrderDataSize"`
	OrderDetailsAvailable w3c.Boolean         `xml:"OrderDetailsAvailable"`
	BankSignature         SignatureType       `xml:"BankSignature,omitempty"`
	SignerInfo            SignerInfoType      `xml:"SignerInfo,omitempty"`

	Any []*w3c.Any
}
