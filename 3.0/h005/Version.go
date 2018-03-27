// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

// Attribute
type Version struct {
	Value *EncryptionVersionType `xml:"Version,attr"`
}

func NewVersion() *Version {
	return new(Version)
}

func (me *Version) SetValue(value *EncryptionVersionType) {
	me.Value = value
}
