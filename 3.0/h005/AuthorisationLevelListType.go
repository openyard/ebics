// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

// ListType
type AuthorisationLevelListType struct {
	Value []*AuthorisationLevelType `xml:",chardata"`
}

func NewAuthorisationLevelListType() *AuthorisationLevelListType {
	return new(AuthorisationLevelListType)
}

func (me *AuthorisationLevelListType) AddValue(value *AuthorisationLevelType) {
	me.Value = append(me.Value, value)
}
