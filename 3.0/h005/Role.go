// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

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
