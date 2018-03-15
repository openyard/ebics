// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type HVDResponseOrderDataType struct {
	DataDigest            *DataDigestType      `xml:"DataDigest"`
	DisplayFile           *w3c.Base64Binary    `xml:"DisplayFile"`
	OrderDataAvailable    *w3c.Boolean         `xml:"OrderDataAvailable"`
	OrderDataSize         *w3c.PositiveInteger `xml:"OrderDataSize"`
	OrderDetailsAvailable *w3c.Boolean         `xml:"OrderDetailsAvailable"`
	BankSignature         *SignatureType       `xml:"BankSignature,omitempty"`
	SignerInfo            []*SignerInfoType    `xml:"SignerInfo,omitempty"`

	Any []*w3c.Any
}
