// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type EbicsKeyManagementResponseBodyDataTransfer struct {
	DataEncryptionInfo *EbicsKeyManagementResponseBodyDataTransferDataEncryptionInfo `xml:"DataEncryptionInfo"`
	OrderData          *EbicsKeyManagementResponseBodyDataTransferOrderData          `xml:"OrderData"`

	Any []*w3c.Any
}

func NewEbicsKeyManagementResponseBodyDataTransfer() *EbicsKeyManagementResponseBodyDataTransfer {
	return new(EbicsKeyManagementResponseBodyDataTransfer)
}

func (me *EbicsKeyManagementResponseBodyDataTransfer) SetDataEncryptionInfo(value *EbicsKeyManagementResponseBodyDataTransferDataEncryptionInfo) {
	me.DataEncryptionInfo = value
}

func (me *EbicsKeyManagementResponseBodyDataTransfer) AddDataEncryptionInfo() *EbicsKeyManagementResponseBodyDataTransferDataEncryptionInfo {
	me.DataEncryptionInfo = new(EbicsKeyManagementResponseBodyDataTransferDataEncryptionInfo)
	return me.DataEncryptionInfo
}

func (me *EbicsKeyManagementResponseBodyDataTransfer) SetOrderData(value *EbicsKeyManagementResponseBodyDataTransferOrderData) {
	me.OrderData = value
}

func (me *EbicsKeyManagementResponseBodyDataTransfer) AddOrderData() *EbicsKeyManagementResponseBodyDataTransferOrderData {
	me.OrderData = new(EbicsKeyManagementResponseBodyDataTransferOrderData)
	return me.OrderData
}
