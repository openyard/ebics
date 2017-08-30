// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
package h005

import ds "github.com/openyard/ebics/3.0/xmldsig"

// complex type
type EbicsRequestBody struct {
	PreValidation   EbicsRequestBodyPreValidation   `xml:"PreValidation,omitempty"`
	DataTransfer    DataTransferRequestType         `xml:"DataTransfer,omitempty"`
	TransferReceipt EbicsRequestBodyTransferReceipt `xml:"TransferReceipt"`
	X509Data        ds.X509Data                     `xml:"http://www.w3.org/2000/09/xmldsig# X509Data,omitempty"`
}
