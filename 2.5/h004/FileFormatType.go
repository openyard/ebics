// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
