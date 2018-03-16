// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

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
