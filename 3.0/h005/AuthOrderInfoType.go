// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type AuthOrderInfoType struct {
	AdminOrderType OrderTBaseType         `xml:"AdminOrderType"`
	Service        RestrictedServiceType  `xml:"Service,omitempty"`
	Description    OrderDescriptionType   `xml:"Description"`
	NumSigRequired w3c.NonNegativeInteger `xml:"NumSigRequired,omitempty"`

	Any []*w3c.Any
}
