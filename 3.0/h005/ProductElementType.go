// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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
