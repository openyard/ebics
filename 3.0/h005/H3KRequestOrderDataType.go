// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package h005

// complex type
type H3KRequestOrderDataType struct {
	SignatureCertificateInfo      SignatureCertificateInfoType      `xml:"SignatureCertificateInfo"`
	AuthenticationCertificateInfo AuthenticationCertificateInfoType `xml:"AuthenticationCertificateInfo"`
	EncryptionCertificateInfo     EncryptionCertificateInfoType     `xml:"EncryptionCertificateInfo"`
	PartnerID                     PartnerIDType                     `xml:"PartnerID"`
	UserID                        UserIDType                        `xml:"UserID"`
}
