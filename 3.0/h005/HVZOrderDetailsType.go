// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
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