// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type UserInfoType struct {
	UserID     UserInfoTypeUserID `xml:"UserID"`
	Name       NameType           `xml:"Name,omitempty"`
	Permission UserPermissionType `xml:"Permission"`

	Any []*w3c.Any
}
