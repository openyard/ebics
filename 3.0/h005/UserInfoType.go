// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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

func (me *UserInfoType) AddUserID() *UserInfoTypeUserID {
	me.UserID = new(UserInfoTypeUserID)
	return me.UserID
}

func (me *UserInfoType) SetName(value *NameType) {
	me.Name = value
}

func (me *UserInfoType) AddName() *NameType {
	me.Name = new(NameType)
	return me.Name
}

func (me *UserInfoType) AddPermission(value *UserPermissionType) {
	me.Permission = append(me.Permission, value)
}
