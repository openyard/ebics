// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// Attribute
type SignatureVersion struct {
	Value *SignatureVersionType `xml:"SignatureVersion,attr"`
}

func NewSignatureVersion() *SignatureVersion {
	return new(SignatureVersion)
}

func (me *SignatureVersion) SetValue(value *SignatureVersionType) {
	me.Value = value
}
