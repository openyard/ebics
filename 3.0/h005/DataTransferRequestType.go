// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type DataTransferRequestType struct {
	DataEncryptionInfo  *DataTransferRequestTypeDataEncryptionInfo `xml:"DataEncryptionInfo"`
	SignatureData       *DataTransferRequestTypeSignatureData      `xml:"SignatureData"`
	DataDigest          *DataDigestType                            `xml:"DataDigest"`
	AdditionalOrderInfo *String255Type                             `xml:"AdditionalOrderInfo,omitempty"`
	OrderData           *DataTransferRequestTypeOrderData          `xml:"OrderData"`

	Any []*w3c.Any
}
