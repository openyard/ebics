// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

// ComplexElement
type StaticHeaderTypeProduct struct {
	Value       *ProductType     `xml:",chardata"`
	Language    *LanguageType    `xml:"Language,attr"`
	InstituteID *InstituteIDType `xml:"InstituteID,attr,omitempty"`
}

func NewStaticHeaderTypeProduct() *StaticHeaderTypeProduct {
	return new(StaticHeaderTypeProduct)
}

func (me *StaticHeaderTypeProduct) SetLanguage(value *LanguageType) {
	me.Language = value
}

func (me *StaticHeaderTypeProduct) SetInstituteID(value *InstituteIDType) {
	me.InstituteID = value
}
