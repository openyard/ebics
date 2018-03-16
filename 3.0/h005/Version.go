// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
