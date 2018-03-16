// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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

func (me *DSAKeyValueType) AddG() *CryptoBinary {
	me.G = new(CryptoBinary)
	return me.G
}

func (me *DSAKeyValueType) SetY(value *CryptoBinary) {
	me.Y = value
}

func (me *DSAKeyValueType) AddY() *CryptoBinary {
	me.Y = new(CryptoBinary)
	return me.Y
}

func (me *DSAKeyValueType) SetJ(value *CryptoBinary) {
	me.J = value
}

func (me *DSAKeyValueType) AddJ() *CryptoBinary {
	me.J = new(CryptoBinary)
	return me.J
}
