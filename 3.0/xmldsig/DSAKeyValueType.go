// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package xmldsig

// ComplexType
type DSAKeyValueType struct {
	G *CryptoBinary `xml:"G,omitempty"`
	Y *CryptoBinary `xml:"Y"`
	J *CryptoBinary `xml:"J,omitempty"`
}

func NewDSAKeyValueType() *DSAKeyValueType {
	return new(DSAKeyValueType)
}

func (me *DSAKeyValueType) SetG(value *CryptoBinary) {
	me.G = value
}

func (me *DSAKeyValueType) SetY(value *CryptoBinary) {
	me.Y = value
}

func (me *DSAKeyValueType) SetJ(value *CryptoBinary) {
	me.J = value
}
