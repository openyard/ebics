// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"
// complex type
type UserInfoType struct {
    
    UserID UserInfoTypeUserID `xml:"UserID"`
    Name NameType `xml:"Name,omitempty"`
    Permission []UserPermissionType `xml:"Permission"`
    
        Any []*w3c.Any
    }
