// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type HVUOrderDetailsType struct {
	OrderType      *OrderTBaseType        `xml:"OrderType"`
	FileFormat     *FileFormatType        `xml:"FileFormat,omitempty"`
	OrderID        *OrderIDType           `xml:"OrderID"`
	OrderDataSize  *w3c.PositiveInteger   `xml:"OrderDataSize"`
	SigningInfo    *HVUSigningInfoType    `xml:"SigningInfo"`
	SignerInfo     []*SignerInfoType      `xml:"SignerInfo,omitempty"`
	OriginatorInfo *HVUOriginatorInfoType `xml:"OriginatorInfo"`

	Any []*w3c.Any
}

func NewHVUOrderDetailsType() *HVUOrderDetailsType {
	return new(HVUOrderDetailsType)
}

func (me *HVUOrderDetailsType) SetOrderType(value *OrderTBaseType) {
	me.OrderType = value
}

func (me *HVUOrderDetailsType) AddOrderType() *OrderTBaseType {
	me.OrderType = new(OrderTBaseType)
	return me.OrderType
}

func (me *HVUOrderDetailsType) SetFileFormat(value *FileFormatType) {
	me.FileFormat = value
}

func (me *HVUOrderDetailsType) AddFileFormat() *FileFormatType {
	me.FileFormat = new(FileFormatType)
	return me.FileFormat
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
