// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type DataTransferResponseType struct {
	DataEncryptionInfo *DataTransferResponseTypeDataEncryptionInfo `xml:"DataEncryptionInfo"`
	SignatureData      *DataTransferResponseTypeSignatureData      `xml:"SignatureData,omitempty"`
	OrderData          *DataTransferResponseTypeOrderData          `xml:"OrderData"`

	Any []*w3c.Any
}

func NewDataTransferResponseType() *DataTransferResponseType {
	return new(DataTransferResponseType)
}

func (me *DataTransferResponseType) SetDataEncryptionInfo(value *DataTransferResponseTypeDataEncryptionInfo) {
	me.DataEncryptionInfo = value
}

func (me *DataTransferResponseType) AddDataEncryptionInfo() *DataTransferResponseTypeDataEncryptionInfo {
	me.DataEncryptionInfo = new(DataTransferResponseTypeDataEncryptionInfo)
	return me.DataEncryptionInfo
}

func (me *DataTransferResponseType) SetSignatureData(value *DataTransferResponseTypeSignatureData) {
	me.SignatureData = value
}

func (me *DataTransferResponseType) AddSignatureData() *DataTransferResponseTypeSignatureData {
	me.SignatureData = new(DataTransferResponseTypeSignatureData)
	return me.SignatureData
}

func (me *DataTransferResponseType) SetOrderData(value *DataTransferResponseTypeOrderData) {
	me.OrderData = value
}

func (me *DataTransferResponseType) AddOrderData() *DataTransferResponseTypeOrderData {
	me.OrderData = new(DataTransferResponseTypeOrderData)
	return me.OrderData
}
