// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

// Attribute
type Description struct {
	Value *AccountDescriptionType `xml:"Description,attr"`
}

func NewDescription() *Description {
	return new(Description)
}

func (me *Description) SetValue(value *AccountDescriptionType) {
	me.Value = value
}
