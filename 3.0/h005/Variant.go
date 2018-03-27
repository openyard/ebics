// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

// Attribute
type Variant struct {
	Value *NumStringType `xml:"variant,attr"`
}

func NewVariant() *Variant {
	return new(Variant)
}

func (me *Variant) SetValue(value *NumStringType) {
	me.Value = value
}
