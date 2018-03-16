// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package s002

import w3c "github.com/openyard/ebics/3.0/w3c"

// SimpleType
type UserIDType w3c.Token

func NewUserIDType(value *w3c.Token) *UserIDType {
	me := (*UserIDType)(value)
	return me
}
