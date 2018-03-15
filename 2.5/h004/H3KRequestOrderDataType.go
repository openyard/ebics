// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h004

// complex type
type H3KRequestOrderDataType struct {
	SignatureCertificateInfo      *SignatureCertificateInfoType      `xml:"SignatureCertificateInfo"`
	AuthenticationCertificateInfo *AuthenticationCertificateInfoType `xml:"AuthenticationCertificateInfo"`
	EncryptionCertificateInfo     *EncryptionCertificateInfoType     `xml:"EncryptionCertificateInfo"`
	PartnerID                     *PartnerIDType                     `xml:"PartnerID"`
	UserID                        *UserIDType                        `xml:"UserID"`
}
