// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type UserInfoType struct {
	UserID     *UserInfoTypeUserID   `xml:"UserID"`
	Name       *NameType             `xml:"Name,omitempty"`
	Permission []*UserPermissionType `xml:"Permission"`

	Any []*w3c.Any
}

func NewUserInfoType() *UserInfoType {
	return new(UserInfoType)
}

func (me *UserInfoType) SetUserID(value *UserInfoTypeUserID) {
	me.UserID = value
}

func (me *UserInfoType) SetName(value *NameType) {
	me.Name = value
}

func (me *UserInfoType) AddPermission(value *UserPermissionType) {
	me.Permission = append(me.Permission, value)
}
