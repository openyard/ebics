// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
