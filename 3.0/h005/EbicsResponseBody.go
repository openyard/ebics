// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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

func (me *EbicsResponseBody) AddDataTransfer() *DataTransferResponseType {
	me.DataTransfer = new(DataTransferResponseType)
	return me.DataTransfer
}

func (me *EbicsResponseBody) SetReturnCode(value *EbicsResponseBodyReturnCode) {
	me.ReturnCode = value
}

func (me *EbicsResponseBody) AddReturnCode() *EbicsResponseBodyReturnCode {
	me.ReturnCode = new(EbicsResponseBodyReturnCode)
	return me.ReturnCode
}

func (me *EbicsResponseBody) SetTimestampBankParameter(value *EbicsResponseBodyTimestampBankParameter) {
	me.TimestampBankParameter = value
}

func (me *EbicsResponseBody) AddTimestampBankParameter() *EbicsResponseBodyTimestampBankParameter {
	me.TimestampBankParameter = new(EbicsResponseBodyTimestampBankParameter)
	return me.TimestampBankParameter
}
