// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

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
