// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

// ComplexType
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

func NewStaticHeaderType() *StaticHeaderType {
	return new(StaticHeaderType)
}

func (me *StaticHeaderType) SetHostID(value *HostIDType) {
	me.HostID = value
}

func (me *StaticHeaderType) SetNonce(value *NonceType) {
	me.Nonce = value
}

func (me *StaticHeaderType) SetTimestamp(value *TimestampType) {
	me.Timestamp = value
}

func (me *StaticHeaderType) SetPartnerID(value *PartnerIDType) {
	me.PartnerID = value
}

func (me *StaticHeaderType) SetUserID(value *UserIDType) {
	me.UserID = value
}

func (me *StaticHeaderType) SetSystemID(value *UserIDType) {
	me.SystemID = value
}

func (me *StaticHeaderType) SetProduct(value *StaticHeaderTypeProduct) {
	me.Product = value
}

func (me *StaticHeaderType) SetOrderDetails(value *StaticHeaderOrderDetailsType) {
	me.OrderDetails = value
}

func (me *StaticHeaderType) SetBankPubKeyDigests(value *StaticHeaderTypeBankPubKeyDigests) {
	me.BankPubKeyDigests = value
}

func (me *StaticHeaderType) SetSecurityMedium(value *SecurityMediumType) {
	me.SecurityMedium = value
}

func (me *StaticHeaderType) SetNumSegments(value *NumSegmentsType) {
	me.NumSegments = value
}

func (me *StaticHeaderType) SetTransactionID(value *TransactionIDType) {
	me.TransactionID = value
}
