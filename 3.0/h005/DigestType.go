// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import ds "github.com/openyard/ebics/3.0/xmldsig"

// SimpleType
type DigestType ds.DigestValueType

func NewDigestType(value *ds.DigestValueType) *DigestType {
	me := (*DigestType)(value)
	return me
}
