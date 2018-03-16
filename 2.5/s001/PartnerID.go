// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package s001

// Attribute
type PartnerID struct {
	Value *string `xml:"PartnerID,attr"`
}

func NewPartnerID() *PartnerID {
	return new(PartnerID)
}

func (me *PartnerID) SetValue(value *string) {
	me.Value = value
}
