// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

// Attribute
type Valid_from struct {
	Value *TimestampType `xml:"valid_from,attr"`
}

func NewValid_from() *Valid_from {
	return new(Valid_from)
}

func (me *Valid_from) SetValue(value *TimestampType) {
	me.Value = value
}
