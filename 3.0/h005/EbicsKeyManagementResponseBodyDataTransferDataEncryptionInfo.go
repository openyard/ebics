// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

// ComplexElement
type EbicsKeyManagementResponseBodyDataTransferDataEncryptionInfo struct {
	Value *DataEncryptionInfoType `xml:",chardata"`

	AuthenticationMarker
}

func NewEbicsKeyManagementResponseBodyDataTransferDataEncryptionInfo() *EbicsKeyManagementResponseBodyDataTransferDataEncryptionInfo {
	return new(EbicsKeyManagementResponseBodyDataTransferDataEncryptionInfo)
}
