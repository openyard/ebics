// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type HVDResponseOrderDataType struct {
	DataDigest            *DataDigestType      `xml:"DataDigest"`
	DisplayFile           *w3c.Base64Binary    `xml:"DisplayFile"`
	OrderDataAvailable    *w3c.Boolean         `xml:"OrderDataAvailable"`
	OrderDataSize         *w3c.PositiveInteger `xml:"OrderDataSize"`
	OrderDetailsAvailable *w3c.Boolean         `xml:"OrderDetailsAvailable"`
	BankSignature         *SignatureType       `xml:"BankSignature,omitempty"`
	SignerInfo            []*SignerInfoType    `xml:"SignerInfo,omitempty"`

	Any []*w3c.Any
}

func NewHVDResponseOrderDataType() *HVDResponseOrderDataType {
	return new(HVDResponseOrderDataType)
}

func (me *HVDResponseOrderDataType) SetDataDigest(value *DataDigestType) {
	me.DataDigest = value
}

func (me *HVDResponseOrderDataType) SetDisplayFile(value *w3c.Base64Binary) {
	me.DisplayFile = value
}

func (me *HVDResponseOrderDataType) SetOrderDataAvailable(value *w3c.Boolean) {
	me.OrderDataAvailable = value
}

func (me *HVDResponseOrderDataType) SetOrderDataSize(value *w3c.PositiveInteger) {
	me.OrderDataSize = value
}

func (me *HVDResponseOrderDataType) SetOrderDetailsAvailable(value *w3c.Boolean) {
	me.OrderDetailsAvailable = value
}

func (me *HVDResponseOrderDataType) SetBankSignature(value *SignatureType) {
	me.BankSignature = value
}

func (me *HVDResponseOrderDataType) AddSignerInfo(value *SignerInfoType) {
	me.SignerInfo = append(me.SignerInfo, value)
}
