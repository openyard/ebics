// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

// ComplexType
type H3KRequestOrderDataType struct {
	SignatureCertificateInfo      *SignatureCertificateInfoType      `xml:"SignatureCertificateInfo"`
	AuthenticationCertificateInfo *AuthenticationCertificateInfoType `xml:"AuthenticationCertificateInfo"`
	EncryptionCertificateInfo     *EncryptionCertificateInfoType     `xml:"EncryptionCertificateInfo"`
	PartnerID                     *PartnerIDType                     `xml:"PartnerID"`
	UserID                        *UserIDType                        `xml:"UserID"`
}

func NewH3KRequestOrderDataType() *H3KRequestOrderDataType {
	return new(H3KRequestOrderDataType)
}

func (me *H3KRequestOrderDataType) SetSignatureCertificateInfo(value *SignatureCertificateInfoType) {
	me.SignatureCertificateInfo = value
}

func (me *H3KRequestOrderDataType) AddSignatureCertificateInfo() *SignatureCertificateInfoType {
	me.SignatureCertificateInfo = new(SignatureCertificateInfoType)
	return me.SignatureCertificateInfo
}

func (me *H3KRequestOrderDataType) SetAuthenticationCertificateInfo(value *AuthenticationCertificateInfoType) {
	me.AuthenticationCertificateInfo = value
}

func (me *H3KRequestOrderDataType) AddAuthenticationCertificateInfo() *AuthenticationCertificateInfoType {
	me.AuthenticationCertificateInfo = new(AuthenticationCertificateInfoType)
	return me.AuthenticationCertificateInfo
}

func (me *H3KRequestOrderDataType) SetEncryptionCertificateInfo(value *EncryptionCertificateInfoType) {
	me.EncryptionCertificateInfo = value
}

func (me *H3KRequestOrderDataType) AddEncryptionCertificateInfo() *EncryptionCertificateInfoType {
	me.EncryptionCertificateInfo = new(EncryptionCertificateInfoType)
	return me.EncryptionCertificateInfo
}

func (me *H3KRequestOrderDataType) SetPartnerID(value *PartnerIDType) {
	me.PartnerID = value
}

func (me *H3KRequestOrderDataType) AddPartnerID() *PartnerIDType {
	me.PartnerID = new(PartnerIDType)
	return me.PartnerID
}

func (me *H3KRequestOrderDataType) SetUserID(value *UserIDType) {
	me.UserID = value
}

func (me *H3KRequestOrderDataType) AddUserID() *UserIDType {
	me.UserID = new(UserIDType)
	return me.UserID
}
