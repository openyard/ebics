// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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

func (me *HVUOrderDetailsType) AddService() *RestrictedServiceType {
	me.Service = new(RestrictedServiceType)
	return me.Service
}

func (me *HVUOrderDetailsType) SetOrderID(value *OrderIDType) {
	me.OrderID = value
}

func (me *HVUOrderDetailsType) AddOrderID() *OrderIDType {
	me.OrderID = new(OrderIDType)
	return me.OrderID
}

func (me *HVUOrderDetailsType) SetOrderDataSize(value *w3c.PositiveInteger) {
	me.OrderDataSize = value
}

func (me *HVUOrderDetailsType) AddOrderDataSize() *w3c.PositiveInteger {
	me.OrderDataSize = new(w3c.PositiveInteger)
	return me.OrderDataSize
}

func (me *HVUOrderDetailsType) SetSigningInfo(value *HVUSigningInfoType) {
	me.SigningInfo = value
}

func (me *HVUOrderDetailsType) AddSigningInfo() *HVUSigningInfoType {
	me.SigningInfo = new(HVUSigningInfoType)
	return me.SigningInfo
}

func (me *HVUOrderDetailsType) AddSignerInfo(value *SignerInfoType) {
	me.SignerInfo = append(me.SignerInfo, value)
}

func (me *HVUOrderDetailsType) SetOriginatorInfo(value *HVUOriginatorInfoType) {
	me.OriginatorInfo = value
}

func (me *HVUOrderDetailsType) AddOriginatorInfo() *HVUOriginatorInfoType {
	me.OriginatorInfo = new(HVUOriginatorInfoType)
	return me.OriginatorInfo
}

func (me *HVUOrderDetailsType) SetAdditionalOrderInfo(value *String255Type) {
	me.AdditionalOrderInfo = value
}

func (me *HVUOrderDetailsType) AddAdditionalOrderInfo() *String255Type {
	me.AdditionalOrderInfo = new(String255Type)
	return me.AdditionalOrderInfo
}
