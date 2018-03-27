// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

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

func (me *PGPDataType) AddPGPKeyID() *w3c.Base64Binary {
	me.PGPKeyID = new(w3c.Base64Binary)
	return me.PGPKeyID
}

func (me *PGPDataType) SetPGPKeyPacket(value *w3c.Base64Binary) {
	me.PGPKeyPacket = value
}

func (me *PGPDataType) AddPGPKeyPacket() *w3c.Base64Binary {
	me.PGPKeyPacket = new(w3c.Base64Binary)
	return me.PGPKeyPacket
}
