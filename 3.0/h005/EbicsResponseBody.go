// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

// ComplexType
type EbicsResponseBody struct {
	DataTransfer           *DataTransferResponseType                `xml:"DataTransfer,omitempty"`
	ReturnCode             *EbicsResponseBodyReturnCode             `xml:"ReturnCode"`
	TimestampBankParameter *EbicsResponseBodyTimestampBankParameter `xml:"TimestampBankParameter,omitempty"`
}

func NewEbicsResponseBody() *EbicsResponseBody {
	return new(EbicsResponseBody)
}

func (me *EbicsResponseBody) SetDataTransfer(value *DataTransferResponseType) {
	me.DataTransfer = value
}

func (me *EbicsResponseBody) SetReturnCode(value *EbicsResponseBodyReturnCode) {
	me.ReturnCode = value
}

func (me *EbicsResponseBody) SetTimestampBankParameter(value *EbicsResponseBodyTimestampBankParameter) {
	me.TimestampBankParameter = value
}
