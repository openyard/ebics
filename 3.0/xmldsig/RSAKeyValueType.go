// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package xmldsig

// ComplexType
type RSAKeyValueType struct {
	Modulus  *CryptoBinary `xml:"Modulus"`
	Exponent *CryptoBinary `xml:"Exponent"`
}

func NewRSAKeyValueType() *RSAKeyValueType {
	return new(RSAKeyValueType)
}

func (me *RSAKeyValueType) SetModulus(value *CryptoBinary) {
	me.Modulus = value
}

func (me *RSAKeyValueType) AddModulus() *CryptoBinary {
	me.Modulus = new(CryptoBinary)
	return me.Modulus
}

func (me *RSAKeyValueType) SetExponent(value *CryptoBinary) {
	me.Exponent = value
}

func (me *RSAKeyValueType) AddExponent() *CryptoBinary {
	me.Exponent = new(CryptoBinary)
	return me.Exponent
}
