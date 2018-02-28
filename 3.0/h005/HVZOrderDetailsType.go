// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type HVZOrderDetailsType struct {
	Service               RestrictedServiceType `xml:"Service"`
	OrderID               OrderIDType           `xml:"OrderID"`
	DataDigest            DataDigestType        `xml:"DataDigest"`
	OrderDataAvailable    w3c.Boolean           `xml:"OrderDataAvailable"`
	OrderDataSize         w3c.PositiveInteger   `xml:"OrderDataSize"`
	OrderDetailsAvailable w3c.Boolean           `xml:"OrderDetailsAvailable"`
	SigningInfo           HVUSigningInfoType    `xml:"SigningInfo"`
	SignerInfo            SignerInfoType        `xml:"SignerInfo,omitempty"`
	OriginatorInfo        HVUOriginatorInfoType `xml:"OriginatorInfo"`
	AdditionalOrderInfo   String255Type         `xml:"AdditionalOrderInfo,omitempty"`

	Any []*w3c.Any
}
