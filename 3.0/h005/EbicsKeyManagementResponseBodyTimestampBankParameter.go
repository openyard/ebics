// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

// ComplexElement
type EbicsKeyManagementResponseBodyTimestampBankParameter struct {
	Value *TimestampType `xml:",chardata"`

	AuthenticationMarker
}

func NewEbicsKeyManagementResponseBodyTimestampBankParameter() *EbicsKeyManagementResponseBodyTimestampBankParameter {
	return new(EbicsKeyManagementResponseBodyTimestampBankParameter)
}
