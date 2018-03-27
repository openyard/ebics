// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// ComplexElement
type EbicsKeyManagementResponseBodyReturnCode struct {
	Value *ReturnCodeType `xml:",chardata"`

	AuthenticationMarker
}

func NewEbicsKeyManagementResponseBodyReturnCode() *EbicsKeyManagementResponseBodyReturnCode {
	return new(EbicsKeyManagementResponseBodyReturnCode)
}
