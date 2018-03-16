// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type StaticHeaderBaseType struct {
	HostID         *HostIDType         `xml:"HostID"`
	Nonce          *NonceType          `xml:"Nonce,omitempty"`
	Timestamp      *TimestampType      `xml:"Timestamp,omitempty"`
	PartnerID      *PartnerIDType      `xml:"PartnerID"`
	UserID         *UserIDType         `xml:"UserID"`
	SystemID       *UserIDType         `xml:"SystemID,omitempty"`
	Product        *ProductElementType `xml:"Product,omitempty"`
	OrderDetails   *OrderDetailsType   `xml:"OrderDetails"`
	SecurityMedium *SecurityMediumType `xml:"SecurityMedium"`

	Any []*w3c.Any
}

func NewStaticHeaderBaseType() *StaticHeaderBaseType {
	return new(StaticHeaderBaseType)
}

func (me *StaticHeaderBaseType) SetHostID(value *HostIDType) {
	me.HostID = value
}

func (me *StaticHeaderBaseType) AddHostID() *HostIDType {
	me.HostID = new(HostIDType)
	return me.HostID
}

func (me *StaticHeaderBaseType) SetNonce(value *NonceType) {
	me.Nonce = value
}

func (me *StaticHeaderBaseType) AddNonce() *NonceType {
	me.Nonce = new(NonceType)
	return me.Nonce
}

func (me *StaticHeaderBaseType) SetTimestamp(value *TimestampType) {
	me.Timestamp = value
}

func (me *StaticHeaderBaseType) AddTimestamp() *TimestampType {
	me.Timestamp = new(TimestampType)
	return me.Timestamp
}

func (me *StaticHeaderBaseType) SetPartnerID(value *PartnerIDType) {
	me.PartnerID = value
}

func (me *StaticHeaderBaseType) AddPartnerID() *PartnerIDType {
	me.PartnerID = new(PartnerIDType)
	return me.PartnerID
}

func (me *StaticHeaderBaseType) SetUserID(value *UserIDType) {
	me.UserID = value
}

func (me *StaticHeaderBaseType) AddUserID() *UserIDType {
	me.UserID = new(UserIDType)
	return me.UserID
}

func (me *StaticHeaderBaseType) SetSystemID(value *UserIDType) {
	me.SystemID = value
}

func (me *StaticHeaderBaseType) AddSystemID() *UserIDType {
	me.SystemID = new(UserIDType)
	return me.SystemID
}

func (me *StaticHeaderBaseType) SetProduct(value *ProductElementType) {
	me.Product = value
}

func (me *StaticHeaderBaseType) AddProduct() *ProductElementType {
	me.Product = new(ProductElementType)
	return me.Product
}

func (me *StaticHeaderBaseType) SetOrderDetails(value *OrderDetailsType) {
	me.OrderDetails = value
}

func (me *StaticHeaderBaseType) AddOrderDetails() *OrderDetailsType {
	me.OrderDetails = new(OrderDetailsType)
	return me.OrderDetails
}

func (me *StaticHeaderBaseType) SetSecurityMedium(value *SecurityMediumType) {
	me.SecurityMedium = value
}

func (me *StaticHeaderBaseType) AddSecurityMedium() *SecurityMediumType {
	me.SecurityMedium = new(SecurityMediumType)
	return me.SecurityMedium
}
