// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package h004

// complex element
type StaticHeaderTypeProduct struct {
	Value       ProductType     `xml:",chardata"`
	Language    LanguageType    `xml:"Language,attr"`
	InstituteID InstituteIDType `xml:"InstituteID,attr,omitempty"`
}
