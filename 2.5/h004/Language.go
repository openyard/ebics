// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// Attribute
type Language struct {
	Value *LanguageType `xml:"Language,attr"`
}

func NewLanguage() *Language {
	return new(Language)
}

func (me *Language) SetValue(value *LanguageType) {
	me.Value = value
}
