// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// ComplexType
type DataDigestType struct {
	DigestType
	SignatureVersion *SignatureVersionType `xml:"SignatureVersion,attr,omitempty"`
}

func NewDataDigestType() *DigestType {
	return new(DigestType)
}

func (me *DataDigestType) SetSignatureVersion(value *SignatureVersionType) {
	me.SignatureVersion = value
}
