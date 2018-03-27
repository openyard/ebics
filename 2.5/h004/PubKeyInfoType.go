// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import ds "github.com/openyard/ebics/2.5/xmldsig"
import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type PubKeyInfoType struct {
	PubKeyValue *PubKeyValueType `xml:"PubKeyValue"`
	X509Data    *ds.X509Data     `xml:"http://www.w3.org/2000/09/xmldsig# X509Data,omitempty"`

	Any []*w3c.Any
}

func NewPubKeyInfoType() *PubKeyInfoType {
	return new(PubKeyInfoType)
}

func (me *PubKeyInfoType) SetPubKeyValue(value *PubKeyValueType) {
	me.PubKeyValue = value
}

func (me *PubKeyInfoType) AddPubKeyValue() *PubKeyValueType {
	me.PubKeyValue = new(PubKeyValueType)
	return me.PubKeyValue
}

func (me *PubKeyInfoType) SetX509Data(value *ds.X509Data) {
	me.X509Data = value
}
