// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *DataTransferRequestType) SetSignatureData(value *DataTransferRequestTypeSignatureData) {
	me.SignatureData = value
}

func (me *DataTransferRequestType) SetDataDigest(value *DataDigestType) {
	me.DataDigest = value
}

func (me *DataTransferRequestType) SetAdditionalOrderInfo(value *String255Type) {
	me.AdditionalOrderInfo = value
}

func (me *DataTransferRequestType) SetOrderData(value *DataTransferRequestTypeOrderData) {
	me.OrderData = value
}
