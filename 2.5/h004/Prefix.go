// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

// Attribute
type Prefix struct {
	Value *BankCodePrefixType `xml:"Prefix,attr"`
}

func NewPrefix() *Prefix {
	return new(Prefix)
}

func (me *Prefix) SetValue(value *BankCodePrefixType) {
	me.Value = value
}
