// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

// Attribute
type CountryCode struct {
	Value *CountryCodeType `xml:"CountryCode,attr"`
}

func NewCountryCode() *CountryCode {
	return new(CountryCode)
}

func (me *CountryCode) SetValue(value *CountryCodeType) {
	me.Value = value
}
