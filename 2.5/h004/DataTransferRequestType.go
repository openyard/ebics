// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type DataTransferRequestType struct {
	DataEncryptionInfo *DataTransferRequestTypeDataEncryptionInfo `xml:"DataEncryptionInfo"`
	SignatureData      *DataTransferRequestTypeSignatureData      `xml:"SignatureData"`
	OrderData          *DataTransferRequestTypeOrderData          `xml:"OrderData"`

	Any []*w3c.Any
}

func NewDataTransferRequestType() *DataTransferRequestType {
	return new(DataTransferRequestType)
}

func (me *DataTransferRequestType) SetDataEncryptionInfo(value *DataTransferRequestTypeDataEncryptionInfo) {
	me.DataEncryptionInfo = value
}

func (me *DataTransferRequestType) SetSignatureData(value *DataTransferRequestTypeSignatureData) {
	me.SignatureData = value
}

func (me *DataTransferRequestType) SetOrderData(value *DataTransferRequestTypeOrderData) {
	me.OrderData = value
}
