// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

// ComplexElement
type PartnerInfoTypeAccountInfo struct {
	Value *AccountType   `xml:",chardata"`
	ID    *AccountIDType `xml:"ID,attr"`
}

func NewPartnerInfoTypeAccountInfo() *PartnerInfoTypeAccountInfo {
	return new(PartnerInfoTypeAccountInfo)
}

func (me *PartnerInfoTypeAccountInfo) SetID(value *AccountIDType) {
	me.ID = value
}
