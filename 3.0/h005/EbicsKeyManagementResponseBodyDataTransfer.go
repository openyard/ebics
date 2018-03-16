// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

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

func (me *EbicsKeyManagementResponseBodyDataTransfer) SetOrderData(value *EbicsKeyManagementResponseBodyDataTransferOrderData) {
	me.OrderData = value
}
