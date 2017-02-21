// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
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
