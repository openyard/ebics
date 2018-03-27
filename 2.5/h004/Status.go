// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// Attribute
type Status struct {
	Value *UserStatusType `xml:"Status,attr"`
}

func NewStatus() *Status {
	return new(Status)
}

func (me *Status) SetValue(value *UserStatusType) {
	me.Value = value
}
