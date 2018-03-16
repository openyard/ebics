// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type PGPDataType struct {
	PGPKeyID     *w3c.Base64Binary `xml:"PGPKeyID"`
	PGPKeyPacket *w3c.Base64Binary `xml:"PGPKeyPacket,omitempty"`

	Any []*w3c.Any
}

func NewPGPDataType() *PGPDataType {
	return new(PGPDataType)
}

func (me *PGPDataType) SetPGPKeyID(value *w3c.Base64Binary) {
	me.PGPKeyID = value
}

func (me *PGPDataType) SetPGPKeyPacket(value *w3c.Base64Binary) {
	me.PGPKeyPacket = value
}
