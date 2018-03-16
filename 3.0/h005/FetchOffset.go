// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
