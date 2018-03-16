// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

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

func (me *StaticHeaderType) AddHostID() *HostIDType {
	me.HostID = new(HostIDType)
	return me.HostID
}

func (me *StaticHeaderType) SetNonce(value *NonceType) {
	me.Nonce = value
}

func (me *StaticHeaderType) AddNonce() *NonceType {
	me.Nonce = new(NonceType)
	return me.Nonce
}

func (me *StaticHeaderType) SetTimestamp(value *TimestampType) {
	me.Timestamp = value
}

func (me *StaticHeaderType) AddTimestamp() *TimestampType {
	me.Timestamp = new(TimestampType)
	return me.Timestamp
}

func (me *StaticHeaderType) SetPartnerID(value *PartnerIDType) {
	me.PartnerID = value
}

func (me *StaticHeaderType) AddPartnerID() *PartnerIDType {
	me.PartnerID = new(PartnerIDType)
	return me.PartnerID
}

func (me *StaticHeaderType) SetUserID(value *UserIDType) {
	me.UserID = value
}

func (me *StaticHeaderType) AddUserID() *UserIDType {
	me.UserID = new(UserIDType)
	return me.UserID
}

func (me *StaticHeaderType) SetSystemID(value *UserIDType) {
	me.SystemID = value
}

func (me *StaticHeaderType) AddSystemID() *UserIDType {
	me.SystemID = new(UserIDType)
	return me.SystemID
}

func (me *StaticHeaderType) SetProduct(value *StaticHeaderTypeProduct) {
	me.Product = value
}

func (me *StaticHeaderType) AddProduct() *StaticHeaderTypeProduct {
	me.Product = new(StaticHeaderTypeProduct)
	return me.Product
}

func (me *StaticHeaderType) SetOrderDetails(value *StaticHeaderOrderDetailsType) {
	me.OrderDetails = value
}

func (me *StaticHeaderType) AddOrderDetails() *StaticHeaderOrderDetailsType {
	me.OrderDetails = new(StaticHeaderOrderDetailsType)
	return me.OrderDetails
}

func (me *StaticHeaderType) SetBankPubKeyDigests(value *StaticHeaderTypeBankPubKeyDigests) {
	me.BankPubKeyDigests = value
}

func (me *StaticHeaderType) AddBankPubKeyDigests() *StaticHeaderTypeBankPubKeyDigests {
	me.BankPubKeyDigests = new(StaticHeaderTypeBankPubKeyDigests)
	return me.BankPubKeyDigests
}

func (me *StaticHeaderType) SetSecurityMedium(value *SecurityMediumType) {
	me.SecurityMedium = value
}

func (me *StaticHeaderType) AddSecurityMedium() *SecurityMediumType {
	me.SecurityMedium = new(SecurityMediumType)
	return me.SecurityMedium
}

func (me *StaticHeaderType) SetNumSegments(value *NumSegmentsType) {
	me.NumSegments = value
}

func (me *StaticHeaderType) AddNumSegments() *NumSegmentsType {
	me.NumSegments = new(NumSegmentsType)
	return me.NumSegments
}

func (me *StaticHeaderType) SetTransactionID(value *TransactionIDType) {
	me.TransactionID = value
}

func (me *StaticHeaderType) AddTransactionID() *TransactionIDType {
	me.TransactionID = new(TransactionIDType)
	return me.TransactionID
}
