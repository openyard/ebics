// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type UserInfoType struct {
	UserID     UserInfoTypeUserID `xml:"UserID"`
	Name       NameType           `xml:"Name,omitempty"`
	Permission UserPermissionType `xml:"Permission"`

	Any []*w3c.Any
}
