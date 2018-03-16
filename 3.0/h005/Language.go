// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

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
