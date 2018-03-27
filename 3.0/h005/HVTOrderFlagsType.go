// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type HVTOrderFlagsType struct {
	CompleteOrderData *w3c.Boolean                  `xml:"completeOrderData,attr"`
	FetchLimit        *HVTOrderFlagsTypeFetchLimit  `xml:"fetchLimit,attr"`
	FetchOffset       *HVTOrderFlagsTypeFetchOffset `xml:"fetchOffset,attr"`
}

func NewHVTOrderFlagsType() *HVTOrderFlagsType {
	return new(HVTOrderFlagsType)
}

func (me *HVTOrderFlagsType) SetCompleteOrderData(value *w3c.Boolean) {
	me.CompleteOrderData = value
}

func (me *HVTOrderFlagsType) SetFetchLimit(value *HVTOrderFlagsTypeFetchLimit) {
	me.FetchLimit = value
}

func (me *HVTOrderFlagsType) SetFetchOffset(value *HVTOrderFlagsTypeFetchOffset) {
	me.FetchOffset = value
}
