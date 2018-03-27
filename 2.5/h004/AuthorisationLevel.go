// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// Attribute
type AuthorisationLevel struct {
	Value *AuthorisationLevelType `xml:"AuthorisationLevel,attr"`
}

func NewAuthorisationLevel() *AuthorisationLevel {
	return new(AuthorisationLevel)
}

func (me *AuthorisationLevel) SetValue(value *AuthorisationLevelType) {
	me.Value = value
}
