// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexElement
type HVTOrderInfoTypeDescription struct {
	Value *w3c.String                      `xml:",chardata"`
	Type  *HVTOrderInfoTypeDescriptionType `xml:"Type,attr"`
}

func NewHVTOrderInfoTypeDescription() *HVTOrderInfoTypeDescription {
	return new(HVTOrderInfoTypeDescription)
}

func (me *HVTOrderInfoTypeDescription) SetType(value *HVTOrderInfoTypeDescriptionType) {
	me.Type = value
}
