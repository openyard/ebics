// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

// Attribute
type FetchOffset struct {
	Value *HVTOrderFlagsTypeFetchOffset `xml:"fetchOffset,attr"`
}

func NewFetchOffset() *FetchOffset {
	return new(FetchOffset)
}

func (me *FetchOffset) SetValue(value *HVTOrderFlagsTypeFetchOffset) {
	me.Value = value
}
