// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

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

func (me *H3KRequestOrderDataType) SetAuthenticationCertificateInfo(value *AuthenticationCertificateInfoType) {
	me.AuthenticationCertificateInfo = value
}

func (me *H3KRequestOrderDataType) SetEncryptionCertificateInfo(value *EncryptionCertificateInfoType) {
	me.EncryptionCertificateInfo = value
}

func (me *H3KRequestOrderDataType) SetPartnerID(value *PartnerIDType) {
	me.PartnerID = value
}

func (me *H3KRequestOrderDataType) SetUserID(value *UserIDType) {
	me.UserID = value
}
