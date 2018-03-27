// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
