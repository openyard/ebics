// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
