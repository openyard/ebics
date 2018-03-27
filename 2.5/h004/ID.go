// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

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
