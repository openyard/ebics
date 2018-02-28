// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package h005

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
