// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type UserPermissionType struct {
	AuthorisationLevel AuthorisationLevelType `xml:"AuthorisationLevel,attr"`
	AdminOrderType     OrderTBaseType         `xml:"AdminOrderType"`
	Service            RestrictedServiceType  `xml:"Service,omitempty"`
	AccountID          AccountIDType          `xml:"AccountID,omitempty"`
	MaxAmount          AmountType             `xml:"MaxAmount,omitempty"`

	Any []*w3c.Any
}
