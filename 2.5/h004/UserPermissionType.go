// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type UserPermissionType struct {
	AuthorisationLevel AuthorisationLevelType `xml:"AuthorisationLevel,attr"`
	OrderTypes         OrderTListType         `xml:"OrderTypes"`
	FileFormat         FileFormatType         `xml:"FileFormat,omitempty"`
	AccountID          AccountIDType          `xml:"AccountID,omitempty"`
	MaxAmount          AmountType             `xml:"MaxAmount,omitempty"`

	Any []*w3c.Any
}
