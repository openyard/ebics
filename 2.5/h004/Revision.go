// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

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
