// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
