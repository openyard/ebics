// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h005

// complex type
type ProductElementType struct {
	ProductType
	Language    LanguageType    `xml:"Language,attr"`
	InstituteID InstituteIDType `xml:"InstituteID,attr,omitempty"`
}
