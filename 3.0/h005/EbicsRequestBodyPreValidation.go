// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

// ComplexElement
type EbicsRequestBodyPreValidation struct {
	Value *PreValidationRequestType `xml:",chardata"`

	AuthenticationMarker
}

func NewEbicsRequestBodyPreValidation() *EbicsRequestBodyPreValidation {
	return new(EbicsRequestBodyPreValidation)
}
