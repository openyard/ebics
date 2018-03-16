// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

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
