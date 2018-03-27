// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

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

func (me *HVDResponseOrderDataType) AddDataDigest() *DataDigestType {
	me.DataDigest = new(DataDigestType)
	return me.DataDigest
}

func (me *HVDResponseOrderDataType) SetDisplayFile(value *w3c.Base64Binary) {
	me.DisplayFile = value
}

func (me *HVDResponseOrderDataType) AddDisplayFile() *w3c.Base64Binary {
	me.DisplayFile = new(w3c.Base64Binary)
	return me.DisplayFile
}

func (me *HVDResponseOrderDataType) SetOrderDataAvailable(value *w3c.Boolean) {
	me.OrderDataAvailable = value
}

func (me *HVDResponseOrderDataType) AddOrderDataAvailable() *w3c.Boolean {
	me.OrderDataAvailable = new(w3c.Boolean)
	return me.OrderDataAvailable
}

func (me *HVDResponseOrderDataType) SetOrderDataSize(value *w3c.PositiveInteger) {
	me.OrderDataSize = value
}

func (me *HVDResponseOrderDataType) AddOrderDataSize() *w3c.PositiveInteger {
	me.OrderDataSize = new(w3c.PositiveInteger)
	return me.OrderDataSize
}

func (me *HVDResponseOrderDataType) SetOrderDetailsAvailable(value *w3c.Boolean) {
	me.OrderDetailsAvailable = value
}

func (me *HVDResponseOrderDataType) AddOrderDetailsAvailable() *w3c.Boolean {
	me.OrderDetailsAvailable = new(w3c.Boolean)
	return me.OrderDetailsAvailable
}

func (me *HVDResponseOrderDataType) SetBankSignature(value *SignatureType) {
	me.BankSignature = value
}

func (me *HVDResponseOrderDataType) AddBankSignature() *SignatureType {
	me.BankSignature = new(SignatureType)
	return me.BankSignature
}

func (me *HVDResponseOrderDataType) AddSignerInfo(value *SignerInfoType) {
	me.SignerInfo = append(me.SignerInfo, value)
}
