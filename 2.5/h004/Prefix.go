// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
