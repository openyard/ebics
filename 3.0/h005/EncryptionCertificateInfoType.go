// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

// ComplexType
type EncryptionCertificateInfoType struct {
	CertificateInfoType
	EncryptionVersion *EncryptionVersionType `xml:"EncryptionVersion"`
}

func NewEncryptionCertificateInfoType() *CertificateInfoType {
	return new(CertificateInfoType)
}

func (me *EncryptionCertificateInfoType) SetEncryptionVersion(value *EncryptionVersionType) {
	me.EncryptionVersion = value
}

func (me *EncryptionCertificateInfoType) AddEncryptionVersion() *EncryptionVersionType {
	me.EncryptionVersion = new(EncryptionVersionType)
	return me.EncryptionVersion
}
