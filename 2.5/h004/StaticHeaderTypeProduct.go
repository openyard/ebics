// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
