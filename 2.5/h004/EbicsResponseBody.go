// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

// complex type
type EbicsResponseBody struct {
	DataTransfer           DataTransferResponseType                `xml:"DataTransfer,omitempty"`
	ReturnCode             EbicsResponseBodyReturnCode             `xml:"ReturnCode"`
	TimestampBankParameter EbicsResponseBodyTimestampBankParameter `xml:"TimestampBankParameter,omitempty"`
}
