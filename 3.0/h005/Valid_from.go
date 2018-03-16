// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
