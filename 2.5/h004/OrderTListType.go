// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
