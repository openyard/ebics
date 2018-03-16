// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// SimpleType
type DisplayFile w3c.Base64Binary

func NewDisplayFile(value *w3c.Base64Binary) *DisplayFile {
	me := (*DisplayFile)(value)
	return me
}
