// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// Attribute
type Role struct {
	Value *AccountNumberRoleType `xml:"Role,attr"`
}

func NewRole() *Role {
	return new(Role)
}

func (me *Role) SetValue(value *AccountNumberRoleType) {
	me.Value = value
}
