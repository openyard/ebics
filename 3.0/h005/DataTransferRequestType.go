// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type DataTransferRequestType struct {
	DataEncryptionInfo  *DataTransferRequestTypeDataEncryptionInfo `xml:"DataEncryptionInfo"`
	SignatureData       *DataTransferRequestTypeSignatureData      `xml:"SignatureData"`
	DataDigest          *DataDigestType                            `xml:"DataDigest"`
	AdditionalOrderInfo *String255Type                             `xml:"AdditionalOrderInfo,omitempty"`
	OrderData           *DataTransferRequestTypeOrderData          `xml:"OrderData"`

	Any []*w3c.Any
}

func NewDataTransferRequestType() *DataTransferRequestType {
	return new(DataTransferRequestType)
}

func (me *DataTransferRequestType) SetDataEncryptionInfo(value *DataTransferRequestTypeDataEncryptionInfo) {
	me.DataEncryptionInfo = value
}

func (me *DataTransferRequestType) AddDataEncryptionInfo() *DataTransferRequestTypeDataEncryptionInfo {
	me.DataEncryptionInfo = new(DataTransferRequestTypeDataEncryptionInfo)
	return me.DataEncryptionInfo
}

func (me *DataTransferRequestType) SetSignatureData(value *DataTransferRequestTypeSignatureData) {
	me.SignatureData = value
}

func (me *DataTransferRequestType) AddSignatureData() *DataTransferRequestTypeSignatureData {
	me.SignatureData = new(DataTransferRequestTypeSignatureData)
	return me.SignatureData
}

func (me *DataTransferRequestType) SetDataDigest(value *DataDigestType) {
	me.DataDigest = value
}

func (me *DataTransferRequestType) AddDataDigest() *DataDigestType {
	me.DataDigest = new(DataDigestType)
	return me.DataDigest
}

func (me *DataTransferRequestType) SetAdditionalOrderInfo(value *String255Type) {
	me.AdditionalOrderInfo = value
}

func (me *DataTransferRequestType) AddAdditionalOrderInfo() *String255Type {
	me.AdditionalOrderInfo = new(String255Type)
	return me.AdditionalOrderInfo
}

func (me *DataTransferRequestType) SetOrderData(value *DataTransferRequestTypeOrderData) {
	me.OrderData = value
}

func (me *DataTransferRequestType) AddOrderData() *DataTransferRequestTypeOrderData {
	me.OrderData = new(DataTransferRequestTypeOrderData)
	return me.OrderData
}
