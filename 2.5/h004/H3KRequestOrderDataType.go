// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h004

// complex type
type H3KRequestOrderDataType struct {
    
    SignatureCertificateInfo SignatureCertificateInfoType `xml:"SignatureCertificateInfo"`
    AuthenticationCertificateInfo AuthenticationCertificateInfoType `xml:"AuthenticationCertificateInfo"`
    EncryptionCertificateInfo EncryptionCertificateInfoType `xml:"EncryptionCertificateInfo"`
    PartnerID PartnerIDType `xml:"PartnerID"`
    UserID UserIDType `xml:"UserID"`
    }
