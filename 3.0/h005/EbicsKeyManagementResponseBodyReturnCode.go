// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

// ComplexElement
type EbicsKeyManagementResponseBodyReturnCode struct {
	Value *ReturnCodeType `xml:",chardata"`

	AuthenticationMarker
}

func NewEbicsKeyManagementResponseBodyReturnCode() *EbicsKeyManagementResponseBodyReturnCode {
	return new(EbicsKeyManagementResponseBodyReturnCode)
}
