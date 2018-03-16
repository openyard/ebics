// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type HVZOrderDetailsType struct {
	Service               *RestrictedServiceType `xml:"Service"`
	OrderID               *OrderIDType           `xml:"OrderID"`
	DataDigest            *DataDigestType        `xml:"DataDigest"`
	OrderDataAvailable    *w3c.Boolean           `xml:"OrderDataAvailable"`
	OrderDataSize         *w3c.PositiveInteger   `xml:"OrderDataSize"`
	OrderDetailsAvailable *w3c.Boolean           `xml:"OrderDetailsAvailable"`
	SigningInfo           *HVUSigningInfoType    `xml:"SigningInfo"`
	SignerInfo            []*SignerInfoType      `xml:"SignerInfo,omitempty"`
	OriginatorInfo        *HVUOriginatorInfoType `xml:"OriginatorInfo"`
	AdditionalOrderInfo   *String255Type         `xml:"AdditionalOrderInfo,omitempty"`

	Any []*w3c.Any
}

func NewHVZOrderDetailsType() *HVZOrderDetailsType {
	return new(HVZOrderDetailsType)
}

func (me *HVZOrderDetailsType) SetService(value *RestrictedServiceType) {
	me.Service = value
}

func (me *HVZOrderDetailsType) AddService() *RestrictedServiceType {
	me.Service = new(RestrictedServiceType)
	return me.Service
}

func (me *HVZOrderDetailsType) SetOrderID(value *OrderIDType) {
	me.OrderID = value
}

func (me *HVZOrderDetailsType) AddOrderID() *OrderIDType {
	me.OrderID = new(OrderIDType)
	return me.OrderID
}

func (me *HVZOrderDetailsType) SetDataDigest(value *DataDigestType) {
	me.DataDigest = value
}

func (me *HVZOrderDetailsType) AddDataDigest() *DataDigestType {
	me.DataDigest = new(DataDigestType)
	return me.DataDigest
}

func (me *HVZOrderDetailsType) SetOrderDataAvailable(value *w3c.Boolean) {
	me.OrderDataAvailable = value
}

func (me *HVZOrderDetailsType) AddOrderDataAvailable() *w3c.Boolean {
	me.OrderDataAvailable = new(w3c.Boolean)
	return me.OrderDataAvailable
}

func (me *HVZOrderDetailsType) SetOrderDataSize(value *w3c.PositiveInteger) {
	me.OrderDataSize = value
}

func (me *HVZOrderDetailsType) AddOrderDataSize() *w3c.PositiveInteger {
	me.OrderDataSize = new(w3c.PositiveInteger)
	return me.OrderDataSize
}

func (me *HVZOrderDetailsType) SetOrderDetailsAvailable(value *w3c.Boolean) {
	me.OrderDetailsAvailable = value
}

func (me *HVZOrderDetailsType) AddOrderDetailsAvailable() *w3c.Boolean {
	me.OrderDetailsAvailable = new(w3c.Boolean)
	return me.OrderDetailsAvailable
}

func (me *HVZOrderDetailsType) SetSigningInfo(value *HVUSigningInfoType) {
	me.SigningInfo = value
}

func (me *HVZOrderDetailsType) AddSigningInfo() *HVUSigningInfoType {
	me.SigningInfo = new(HVUSigningInfoType)
	return me.SigningInfo
}

func (me *HVZOrderDetailsType) AddSignerInfo(value *SignerInfoType) {
	me.SignerInfo = append(me.SignerInfo, value)
}

func (me *HVZOrderDetailsType) SetOriginatorInfo(value *HVUOriginatorInfoType) {
	me.OriginatorInfo = value
}

func (me *HVZOrderDetailsType) AddOriginatorInfo() *HVUOriginatorInfoType {
	me.OriginatorInfo = new(HVUOriginatorInfoType)
	return me.OriginatorInfo
}

func (me *HVZOrderDetailsType) SetAdditionalOrderInfo(value *String255Type) {
	me.AdditionalOrderInfo = value
}

func (me *HVZOrderDetailsType) AddAdditionalOrderInfo() *String255Type {
	me.AdditionalOrderInfo = new(String255Type)
	return me.AdditionalOrderInfo
}
