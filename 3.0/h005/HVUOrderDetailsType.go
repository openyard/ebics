// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type HVUOrderDetailsType struct {
	Service             *RestrictedServiceType `xml:"Service"`
	OrderID             *OrderIDType           `xml:"OrderID"`
	OrderDataSize       *w3c.PositiveInteger   `xml:"OrderDataSize"`
	SigningInfo         *HVUSigningInfoType    `xml:"SigningInfo"`
	SignerInfo          []*SignerInfoType      `xml:"SignerInfo,omitempty"`
	OriginatorInfo      *HVUOriginatorInfoType `xml:"OriginatorInfo"`
	AdditionalOrderInfo *String255Type         `xml:"AdditionalOrderInfo,omitempty"`

	Any []*w3c.Any
}

func NewHVUOrderDetailsType() *HVUOrderDetailsType {
	return new(HVUOrderDetailsType)
}

func (me *HVUOrderDetailsType) SetService(value *RestrictedServiceType) {
	me.Service = value
}

func (me *HVUOrderDetailsType) SetOrderID(value *OrderIDType) {
	me.OrderID = value
}

func (me *HVUOrderDetailsType) SetOrderDataSize(value *w3c.PositiveInteger) {
	me.OrderDataSize = value
}

func (me *HVUOrderDetailsType) SetSigningInfo(value *HVUSigningInfoType) {
	me.SigningInfo = value
}

func (me *HVUOrderDetailsType) AddSignerInfo(value *SignerInfoType) {
	me.SignerInfo = append(me.SignerInfo, value)
}

func (me *HVUOrderDetailsType) SetOriginatorInfo(value *HVUOriginatorInfoType) {
	me.OriginatorInfo = value
}

func (me *HVUOrderDetailsType) SetAdditionalOrderInfo(value *String255Type) {
	me.AdditionalOrderInfo = value
}
