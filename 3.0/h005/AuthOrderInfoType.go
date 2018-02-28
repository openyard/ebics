// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
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
