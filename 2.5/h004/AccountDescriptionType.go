// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// SimpleType
type AccountDescriptionType w3c.NormalizedString

func NewAccountDescriptionType(value *w3c.NormalizedString) *AccountDescriptionType {
	me := (*AccountDescriptionType)(value)
	return me
}
