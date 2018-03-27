// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

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
