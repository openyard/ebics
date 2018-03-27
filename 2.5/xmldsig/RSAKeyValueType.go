// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
