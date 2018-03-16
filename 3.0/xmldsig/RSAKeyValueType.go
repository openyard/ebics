// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *RSAKeyValueType) SetExponent(value *CryptoBinary) {
	me.Exponent = value
}
