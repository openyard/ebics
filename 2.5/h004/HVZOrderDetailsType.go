// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type HVZOrderDetailsType struct {
	OrderType             *OrderTBaseType        `xml:"OrderType"`
	FileFormat            *FileFormatType        `xml:"FileFormat,omitempty"`
	OrderID               *OrderIDType           `xml:"OrderID"`
	DataDigest            *DataDigestType        `xml:"DataDigest"`
	OrderDataAvailable    *w3c.Boolean           `xml:"OrderDataAvailable"`
	OrderDataSize         *w3c.PositiveInteger   `xml:"OrderDataSize"`
	OrderDetailsAvailable *w3c.Boolean           `xml:"OrderDetailsAvailable"`
	SigningInfo           *HVUSigningInfoType    `xml:"SigningInfo"`
	SignerInfo            []*SignerInfoType      `xml:"SignerInfo,omitempty"`
	OriginatorInfo        *HVUOriginatorInfoType `xml:"OriginatorInfo"`

	Any []*w3c.Any
}

func NewHVZOrderDetailsType() *HVZOrderDetailsType {
	return new(HVZOrderDetailsType)
}

func (me *HVZOrderDetailsType) SetOrderType(value *OrderTBaseType) {
	me.OrderType = value
}

func (me *HVZOrderDetailsType) SetFileFormat(value *FileFormatType) {
	me.FileFormat = value
}

func (me *HVZOrderDetailsType) SetOrderID(value *OrderIDType) {
	me.OrderID = value
}

func (me *HVZOrderDetailsType) SetDataDigest(value *DataDigestType) {
	me.DataDigest = value
}

func (me *HVZOrderDetailsType) SetOrderDataAvailable(value *w3c.Boolean) {
	me.OrderDataAvailable = value
}

func (me *HVZOrderDetailsType) SetOrderDataSize(value *w3c.PositiveInteger) {
	me.OrderDataSize = value
}

func (me *HVZOrderDetailsType) SetOrderDetailsAvailable(value *w3c.Boolean) {
	me.OrderDetailsAvailable = value
}

func (me *HVZOrderDetailsType) SetSigningInfo(value *HVUSigningInfoType) {
	me.SigningInfo = value
}

func (me *HVZOrderDetailsType) AddSignerInfo(value *SignerInfoType) {
	me.SignerInfo = append(me.SignerInfo, value)
}

func (me *HVZOrderDetailsType) SetOriginatorInfo(value *HVUOriginatorInfoType) {
	me.OriginatorInfo = value
}
