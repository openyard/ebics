// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

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
