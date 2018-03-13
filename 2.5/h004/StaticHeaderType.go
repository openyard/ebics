// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h004

// complex type
type StaticHeaderType struct {
    
    HostID HostIDType `xml:"HostID"`
    Nonce NonceType `xml:"Nonce"`
    Timestamp TimestampType `xml:"Timestamp"`
    PartnerID PartnerIDType `xml:"PartnerID"`
    UserID UserIDType `xml:"UserID"`
    SystemID UserIDType `xml:"SystemID,omitempty"`
    Product StaticHeaderTypeProduct `xml:"Product,omitempty"`
    OrderDetails StaticHeaderOrderDetailsType `xml:"OrderDetails"`
    BankPubKeyDigests StaticHeaderTypeBankPubKeyDigests `xml:"BankPubKeyDigests"`
    SecurityMedium SecurityMediumType `xml:"SecurityMedium"`
    NumSegments NumSegmentsType `xml:"NumSegments,omitempty"`
    TransactionID TransactionIDType `xml:"TransactionID"`
    }
