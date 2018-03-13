// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h004

import ds "github.com/openyard/ebics/2.5/xmldsig"
// complex type
type EbicsRequestBody struct {
    
    PreValidation EbicsRequestBodyPreValidation `xml:"PreValidation,omitempty"`
    DataTransfer DataTransferRequestType `xml:"DataTransfer,omitempty"`
    TransferReceipt EbicsRequestBodyTransferReceipt `xml:"TransferReceipt"`
    X509Data ds.X509Data `xml:"http://www.w3.org/2000/09/xmldsig# X509Data,omitempty"`
    }
