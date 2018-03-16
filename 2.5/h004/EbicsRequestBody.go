// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *EbicsRequestBody) SetDataTransfer(value *DataTransferRequestType) {
	me.DataTransfer = value
}

func (me *EbicsRequestBody) SetTransferReceipt(value *EbicsRequestBodyTransferReceipt) {
	me.TransferReceipt = value
}

func (me *EbicsRequestBody) SetX509Data(value *ds.X509Data) {
	me.X509Data = value
}
