// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

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

func (me *StaticHeaderBaseType) SetNonce(value *NonceType) {
	me.Nonce = value
}

func (me *StaticHeaderBaseType) SetTimestamp(value *TimestampType) {
	me.Timestamp = value
}

func (me *StaticHeaderBaseType) SetPartnerID(value *PartnerIDType) {
	me.PartnerID = value
}

func (me *StaticHeaderBaseType) SetUserID(value *UserIDType) {
	me.UserID = value
}

func (me *StaticHeaderBaseType) SetSystemID(value *UserIDType) {
	me.SystemID = value
}

func (me *StaticHeaderBaseType) SetProduct(value *ProductElementType) {
	me.Product = value
}

func (me *StaticHeaderBaseType) SetOrderDetails(value *OrderDetailsType) {
	me.OrderDetails = value
}

func (me *StaticHeaderBaseType) SetSecurityMedium(value *SecurityMediumType) {
	me.SecurityMedium = value
}
