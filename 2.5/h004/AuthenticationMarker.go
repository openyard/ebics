// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// AttributeGroup
type AuthenticationMarker struct {
	Authenticate *w3c.Boolean `xml:"authenticate,attr"`
}

func NewAuthenticationMarker() *AuthenticationMarker {
	return new(AuthenticationMarker)
}

func (me *AuthenticationMarker) SetAuthenticate(value *w3c.Boolean) {
	me.Authenticate = value
}
