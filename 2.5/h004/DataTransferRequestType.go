// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type DataTransferRequestType struct {
	DataEncryptionInfo DataTransferRequestTypeDataEncryptionInfo `xml:"DataEncryptionInfo"`
	SignatureData      DataTransferRequestTypeSignatureData      `xml:"SignatureData"`
	OrderData          DataTransferRequestTypeOrderData          `xml:"OrderData"`

	Any []*w3c.Any
}
