// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
