// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

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
