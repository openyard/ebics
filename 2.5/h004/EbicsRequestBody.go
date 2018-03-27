// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import ds "github.com/openyard/ebics/2.5/xmldsig"

// ComplexType
type EbicsRequestBody struct {
	PreValidation   *EbicsRequestBodyPreValidation   `xml:"PreValidation,omitempty"`
	DataTransfer    *DataTransferRequestType         `xml:"DataTransfer,omitempty"`
	TransferReceipt *EbicsRequestBodyTransferReceipt `xml:"TransferReceipt"`
	X509Data        *ds.X509Data                     `xml:"http://www.w3.org/2000/09/xmldsig# X509Data,omitempty"`
}

func NewEbicsRequestBody() *EbicsRequestBody {
	return new(EbicsRequestBody)
}

func (me *EbicsRequestBody) SetPreValidation(value *EbicsRequestBodyPreValidation) {
	me.PreValidation = value
}

func (me *EbicsRequestBody) AddPreValidation() *EbicsRequestBodyPreValidation {
	me.PreValidation = new(EbicsRequestBodyPreValidation)
	return me.PreValidation
}

func (me *EbicsRequestBody) SetDataTransfer(value *DataTransferRequestType) {
	me.DataTransfer = value
}

func (me *EbicsRequestBody) AddDataTransfer() *DataTransferRequestType {
	me.DataTransfer = new(DataTransferRequestType)
	return me.DataTransfer
}

func (me *EbicsRequestBody) SetTransferReceipt(value *EbicsRequestBodyTransferReceipt) {
	me.TransferReceipt = value
}

func (me *EbicsRequestBody) AddTransferReceipt() *EbicsRequestBodyTransferReceipt {
	me.TransferReceipt = new(EbicsRequestBodyTransferReceipt)
	return me.TransferReceipt
}

func (me *EbicsRequestBody) SetX509Data(value *ds.X509Data) {
	me.X509Data = value
}
