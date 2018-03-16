// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *HVUOrderDetailsType) SetFileFormat(value *FileFormatType) {
	me.FileFormat = value
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
