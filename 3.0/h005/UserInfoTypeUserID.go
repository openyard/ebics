// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

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
