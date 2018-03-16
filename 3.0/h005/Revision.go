// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

// Attribute
type Revision struct {
	Value *ProtocolRevisionType `xml:"Revision,attr"`
}

func NewRevision() *Revision {
	return new(Revision)
}

func (me *Revision) SetValue(value *ProtocolRevisionType) {
	me.Value = value
}
