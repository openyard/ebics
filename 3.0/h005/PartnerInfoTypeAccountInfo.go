// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
