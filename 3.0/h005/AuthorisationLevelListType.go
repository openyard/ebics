// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
