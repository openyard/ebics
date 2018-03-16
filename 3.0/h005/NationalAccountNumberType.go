// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// SimpleType
type NationalAccountNumberType w3c.NormalizedString

func NewNationalAccountNumberType(value *w3c.NormalizedString) *NationalAccountNumberType {
	me := (*NationalAccountNumberType)(value)
	return me
}
