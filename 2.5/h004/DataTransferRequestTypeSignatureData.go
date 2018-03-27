// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// ComplexElement
type DataTransferRequestTypeSignatureData struct {
	Value *SignatureDataType `xml:",chardata"`

	AuthenticationMarker
}

func NewDataTransferRequestTypeSignatureData() *DataTransferRequestTypeSignatureData {
	return new(DataTransferRequestTypeSignatureData)
}
