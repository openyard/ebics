// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// ComplexType
type SignatureCertificateInfoType struct {
	CertificateInfoType
	SignatureVersion *SignatureVersionType `xml:"SignatureVersion"`
}

func NewSignatureCertificateInfoType() *CertificateInfoType {
	return new(CertificateInfoType)
}

func (me *SignatureCertificateInfoType) SetSignatureVersion(value *SignatureVersionType) {
	me.SignatureVersion = value
}

func (me *SignatureCertificateInfoType) AddSignatureVersion() *SignatureVersionType {
	me.SignatureVersion = new(SignatureVersionType)
	return me.SignatureVersion
}
