// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

// ComplexType
type EbicsKeyManagementResponseBody struct {
	DataTransfer           *EbicsKeyManagementResponseBodyDataTransfer           `xml:"DataTransfer,omitempty"`
	ReturnCode             *EbicsKeyManagementResponseBodyReturnCode             `xml:"ReturnCode"`
	TimestampBankParameter *EbicsKeyManagementResponseBodyTimestampBankParameter `xml:"TimestampBankParameter,omitempty"`
}

func NewEbicsKeyManagementResponseBody() *EbicsKeyManagementResponseBody {
	return new(EbicsKeyManagementResponseBody)
}

func (me *EbicsKeyManagementResponseBody) SetDataTransfer(value *EbicsKeyManagementResponseBodyDataTransfer) {
	me.DataTransfer = value
}

func (me *EbicsKeyManagementResponseBody) SetReturnCode(value *EbicsKeyManagementResponseBodyReturnCode) {
	me.ReturnCode = value
}

func (me *EbicsKeyManagementResponseBody) SetTimestampBankParameter(value *EbicsKeyManagementResponseBodyTimestampBankParameter) {
	me.TimestampBankParameter = value
}
