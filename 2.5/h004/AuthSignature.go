// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import ds "github.com/openyard/ebics/2.5/xmldsig"

// SimpleType
type AuthSignature ds.SignatureType

func NewAuthSignature(value *ds.SignatureType) *AuthSignature {
	me := (*AuthSignature)(value)
	return me
}
