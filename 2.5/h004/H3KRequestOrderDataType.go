// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

// complex type
type H3KRequestOrderDataType struct {
	SignatureCertificateInfo      SignatureCertificateInfoType      `xml:"SignatureCertificateInfo"`
	AuthenticationCertificateInfo AuthenticationCertificateInfoType `xml:"AuthenticationCertificateInfo"`
	EncryptionCertificateInfo     EncryptionCertificateInfoType     `xml:"EncryptionCertificateInfo"`
	PartnerID                     PartnerIDType                     `xml:"PartnerID"`
	UserID                        UserIDType                        `xml:"UserID"`
}
