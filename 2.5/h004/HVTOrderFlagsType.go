// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

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
