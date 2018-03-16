// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type DataTransferResponseType struct {
	OrderData *DataTransferResponseTypeOrderData `xml:"OrderData"`

	Any []*w3c.Any
}

func NewDataTransferResponseType() *DataTransferResponseType {
	return new(DataTransferResponseType)
}

func (me *DataTransferResponseType) SetOrderData(value *DataTransferResponseTypeOrderData) {
	me.OrderData = value
}

func (me *DataTransferResponseType) AddOrderData() *DataTransferResponseTypeOrderData {
	me.OrderData = new(DataTransferResponseTypeOrderData)
	return me.OrderData
}
