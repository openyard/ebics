// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

// AttributeGroup
type SignerPermission struct {
	AuthorisationLevel *AuthorisationLevelType `xml:"AuthorisationLevel,attr"`
}

func NewSignerPermission() *SignerPermission {
	return new(SignerPermission)
}

func (me *SignerPermission) SetAuthorisationLevel(value *AuthorisationLevelType) {
	me.AuthorisationLevel = value
}
