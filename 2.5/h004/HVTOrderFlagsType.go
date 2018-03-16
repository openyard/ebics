// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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
