// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

// ComplexType
type ProductElementType struct {
	ProductType
	Language    *LanguageType    `xml:"Language,attr"`
	InstituteID *InstituteIDType `xml:"InstituteID,attr,omitempty"`
}

func NewProductElementType() *ProductType {
	return new(ProductType)
}

func (me *ProductElementType) SetLanguage(value *LanguageType) {
	me.Language = value
}

func (me *ProductElementType) SetInstituteID(value *InstituteIDType) {
	me.InstituteID = value
}
