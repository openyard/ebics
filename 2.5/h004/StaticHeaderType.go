// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h004

// complex type
type StaticHeaderType struct {
	HostID            *HostIDType                        `xml:"HostID"`
	Nonce             *NonceType                         `xml:"Nonce"`
	Timestamp         *TimestampType                     `xml:"Timestamp"`
	PartnerID         *PartnerIDType                     `xml:"PartnerID"`
	UserID            *UserIDType                        `xml:"UserID"`
	SystemID          *UserIDType                        `xml:"SystemID,omitempty"`
	Product           *StaticHeaderTypeProduct           `xml:"Product,omitempty"`
	OrderDetails      *StaticHeaderOrderDetailsType      `xml:"OrderDetails"`
	BankPubKeyDigests *StaticHeaderTypeBankPubKeyDigests `xml:"BankPubKeyDigests"`
	SecurityMedium    *SecurityMediumType                `xml:"SecurityMedium"`
	NumSegments       *NumSegmentsType                   `xml:"NumSegments,omitempty"`
	TransactionID     *TransactionIDType                 `xml:"TransactionID"`
}
