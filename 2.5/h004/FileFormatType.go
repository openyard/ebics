// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type FileFormatType struct {
	w3c.Token
	CountryCode *CountryCodeType `xml:"CountryCode,attr"`
}

func NewFileFormatType() *w3c.Token {
	return new(w3c.Token)
}

func (me *FileFormatType) SetCountryCode(value *CountryCodeType) {
	me.CountryCode = value
}
