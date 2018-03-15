// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h005

import ds "github.com/openyard/ebics/3.0/xmldsig"

// complex type
type EbicsRequestBody struct {
	PreValidation   *EbicsRequestBodyPreValidation   `xml:"PreValidation,omitempty"`
	DataTransfer    *DataTransferRequestType         `xml:"DataTransfer,omitempty"`
	TransferReceipt *EbicsRequestBodyTransferReceipt `xml:"TransferReceipt"`
	X509Data        *ds.X509Data                     `xml:"http://www.w3.org/2000/09/xmldsig# X509Data,omitempty"`
}
