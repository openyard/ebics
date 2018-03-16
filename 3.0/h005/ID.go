// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

// Attribute
type ID struct {
	Value *AccountIDType `xml:"ID,attr"`
}

func NewID() *ID {
	return new(ID)
}

func (me *ID) SetValue(value *AccountIDType) {
	me.Value = value
}
