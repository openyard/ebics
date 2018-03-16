// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

// ComplexElement
type UserInfoTypeUserID struct {
	Value  *UserIDType     `xml:",chardata"`
	Status *UserStatusType `xml:"Status,attr"`
}

func NewUserInfoTypeUserID() *UserInfoTypeUserID {
	return new(UserInfoTypeUserID)
}

func (me *UserInfoTypeUserID) SetStatus(value *UserStatusType) {
	me.Status = value
}
