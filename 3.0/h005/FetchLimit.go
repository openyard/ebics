// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

// Attribute
type FetchLimit struct {
	Value *HVTOrderFlagsTypeFetchLimit `xml:"fetchLimit,attr"`
}

func NewFetchLimit() *FetchLimit {
	return new(FetchLimit)
}

func (me *FetchLimit) SetValue(value *HVTOrderFlagsTypeFetchLimit) {
	me.Value = value
}
