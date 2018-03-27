// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

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
