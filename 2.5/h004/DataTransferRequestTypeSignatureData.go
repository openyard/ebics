// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

// ComplexElement
type DataTransferRequestTypeSignatureData struct {
	Value *SignatureDataType `xml:",chardata"`

	AuthenticationMarker
}

func NewDataTransferRequestTypeSignatureData() *DataTransferRequestTypeSignatureData {
	return new(DataTransferRequestTypeSignatureData)
}
