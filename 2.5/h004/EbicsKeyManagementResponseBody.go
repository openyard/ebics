// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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

func (me *EbicsKeyManagementResponseBody) AddDataTransfer() *EbicsKeyManagementResponseBodyDataTransfer {
	me.DataTransfer = new(EbicsKeyManagementResponseBodyDataTransfer)
	return me.DataTransfer
}

func (me *EbicsKeyManagementResponseBody) SetReturnCode(value *EbicsKeyManagementResponseBodyReturnCode) {
	me.ReturnCode = value
}

func (me *EbicsKeyManagementResponseBody) AddReturnCode() *EbicsKeyManagementResponseBodyReturnCode {
	me.ReturnCode = new(EbicsKeyManagementResponseBodyReturnCode)
	return me.ReturnCode
}

func (me *EbicsKeyManagementResponseBody) SetTimestampBankParameter(value *EbicsKeyManagementResponseBodyTimestampBankParameter) {
	me.TimestampBankParameter = value
}

func (me *EbicsKeyManagementResponseBody) AddTimestampBankParameter() *EbicsKeyManagementResponseBodyTimestampBankParameter {
	me.TimestampBankParameter = new(EbicsKeyManagementResponseBodyTimestampBankParameter)
	return me.TimestampBankParameter
}
