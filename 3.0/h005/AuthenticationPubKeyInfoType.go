// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

// ComplexType
type AuthenticationPubKeyInfoType struct {
	PubKeyInfoType
	AuthenticationVersion *AuthenticationVersionType `xml:"AuthenticationVersion"`
}

func NewAuthenticationPubKeyInfoType() *PubKeyInfoType {
	return new(PubKeyInfoType)
}

func (me *AuthenticationPubKeyInfoType) SetAuthenticationVersion(value *AuthenticationVersionType) {
	me.AuthenticationVersion = value
}

func (me *AuthenticationPubKeyInfoType) AddAuthenticationVersion() *AuthenticationVersionType {
	me.AuthenticationVersion = new(AuthenticationVersionType)
	return me.AuthenticationVersion
}
