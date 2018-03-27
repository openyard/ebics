// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package xmldsig

// ComplexType
type DSAKeyValueType struct {
	P           *CryptoBinary `xml:"P"`
	Q           *CryptoBinary `xml:"Q"`
	Seed        *CryptoBinary `xml:"Seed"`
	PgenCounter *CryptoBinary `xml:"PgenCounter"`
	G           *CryptoBinary `xml:"G,omitempty"`
	Y           *CryptoBinary `xml:"Y"`
	J           *CryptoBinary `xml:"J,omitempty"`
}

func NewDSAKeyValueType() *DSAKeyValueType {
	return new(DSAKeyValueType)
}

func (me *DSAKeyValueType) SetP(value *CryptoBinary) {
	me.P = value
}

func (me *DSAKeyValueType) AddP() *CryptoBinary {
	me.P = new(CryptoBinary)
	return me.P
}

func (me *DSAKeyValueType) SetQ(value *CryptoBinary) {
	me.Q = value
}

func (me *DSAKeyValueType) AddQ() *CryptoBinary {
	me.Q = new(CryptoBinary)
	return me.Q
}

func (me *DSAKeyValueType) SetSeed(value *CryptoBinary) {
	me.Seed = value
}

func (me *DSAKeyValueType) AddSeed() *CryptoBinary {
	me.Seed = new(CryptoBinary)
	return me.Seed
}

func (me *DSAKeyValueType) SetPgenCounter(value *CryptoBinary) {
	me.PgenCounter = value
}

func (me *DSAKeyValueType) AddPgenCounter() *CryptoBinary {
	me.PgenCounter = new(CryptoBinary)
	return me.PgenCounter
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
