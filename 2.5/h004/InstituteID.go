// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// Attribute
type InstituteID struct {
	Value *InstituteIDType `xml:"InstituteID,attr"`
}

func NewInstituteID() *InstituteID {
	return new(InstituteID)
}

func (me *InstituteID) SetValue(value *InstituteIDType) {
	me.Value = value
}
