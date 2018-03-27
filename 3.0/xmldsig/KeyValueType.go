// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
