// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
