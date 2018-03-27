// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

// ComplexElement
type DataTransferResponseTypeDataEncryptionInfo struct {
	Value *DataEncryptionInfoType `xml:",chardata"`

	AuthenticationMarker
}

func NewDataTransferResponseTypeDataEncryptionInfo() *DataTransferResponseTypeDataEncryptionInfo {
	return new(DataTransferResponseTypeDataEncryptionInfo)
}
