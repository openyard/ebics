// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type KeyValueType struct {
	DSAKeyValue *DSAKeyValue `xml:"DSAKeyValue"`
	RSAKeyValue *RSAKeyValue `xml:"RSAKeyValue"`

	Any []*w3c.Any
}

func NewKeyValueType() *KeyValueType {
	return new(KeyValueType)
}

func (me *KeyValueType) SetDSAKeyValue(value *DSAKeyValue) {
	me.DSAKeyValue = value
}

func (me *KeyValueType) SetRSAKeyValue(value *RSAKeyValue) {
	me.RSAKeyValue = value
}
