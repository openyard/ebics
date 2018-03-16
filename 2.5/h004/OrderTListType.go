// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

// ListType
type OrderTListType struct {
	Value []*OrderTBaseType `xml:",chardata"`
}

func NewOrderTListType() *OrderTListType {
	return new(OrderTListType)
}

func (me *OrderTListType) AddValue(value *OrderTBaseType) {
	me.Value = append(me.Value, value)
}
