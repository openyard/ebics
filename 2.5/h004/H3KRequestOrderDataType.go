// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
package h004

// complex type
type H3KRequestOrderDataType struct {
	SignatureCertificateInfo      SignatureCertificateInfoType      `xml:"SignatureCertificateInfo"`
	AuthenticationCertificateInfo AuthenticationCertificateInfoType `xml:"AuthenticationCertificateInfo"`
	EncryptionCertificateInfo     EncryptionCertificateInfoType     `xml:"EncryptionCertificateInfo"`
	PartnerID                     PartnerIDType                     `xml:"PartnerID"`
	UserID                        UserIDType                        `xml:"UserID"`
}
