// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

// complex type
type StaticHeaderType struct {
	HostID            HostIDType                        `xml:"HostID"`
	Nonce             NonceType                         `xml:"Nonce"`
	Timestamp         TimestampType                     `xml:"Timestamp"`
	PartnerID         PartnerIDType                     `xml:"PartnerID"`
	UserID            UserIDType                        `xml:"UserID"`
	SystemID          UserIDType                        `xml:"SystemID,omitempty"`
	Product           StaticHeaderTypeProduct           `xml:"Product,omitempty"`
	OrderDetails      StaticHeaderOrderDetailsType      `xml:"OrderDetails"`
	BankPubKeyDigests StaticHeaderTypeBankPubKeyDigests `xml:"BankPubKeyDigests"`
	SecurityMedium    SecurityMediumType                `xml:"SecurityMedium"`
	NumSegments       NumSegmentsType                   `xml:"NumSegments,omitempty"`
	TransactionID     TransactionIDType                 `xml:"TransactionID"`
}
