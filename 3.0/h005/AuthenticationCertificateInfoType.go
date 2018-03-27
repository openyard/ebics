// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

// ComplexType
type AuthenticationCertificateInfoType struct {
	CertificateInfoType
	AuthenticationVersion *AuthenticationVersionType `xml:"AuthenticationVersion"`
}

func NewAuthenticationCertificateInfoType() *CertificateInfoType {
	return new(CertificateInfoType)
}

func (me *AuthenticationCertificateInfoType) SetAuthenticationVersion(value *AuthenticationVersionType) {
	me.AuthenticationVersion = value
}

func (me *AuthenticationCertificateInfoType) AddAuthenticationVersion() *AuthenticationVersionType {
	me.AuthenticationVersion = new(AuthenticationVersionType)
	return me.AuthenticationVersion
}
