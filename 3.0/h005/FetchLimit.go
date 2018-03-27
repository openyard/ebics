// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
