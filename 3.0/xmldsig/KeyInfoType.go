// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type KeyInfoType struct {
	Id              *w3c.ID          `xml:"Id,attr,omitempty"`
	KeyName         *KeyName         `xml:"KeyName"`
	KeyValue        *KeyValue        `xml:"KeyValue"`
	RetrievalMethod *RetrievalMethod `xml:"RetrievalMethod"`
	X509Data        *X509Data        `xml:"X509Data"`
	PGPData         *PGPData         `xml:"PGPData"`
	SPKIData        *SPKIData        `xml:"SPKIData"`
	MgmtData        *MgmtData        `xml:"MgmtData"`

	Any []*w3c.Any
}

func NewKeyInfoType() *KeyInfoType {
	return new(KeyInfoType)
}

func (me *KeyInfoType) SetId(value *w3c.ID) {
	me.Id = value
}

func (me *KeyInfoType) SetKeyName(value *KeyName) {
	me.KeyName = value
}

func (me *KeyInfoType) SetKeyValue(value *KeyValue) {
	me.KeyValue = value
}

func (me *KeyInfoType) SetRetrievalMethod(value *RetrievalMethod) {
	me.RetrievalMethod = value
}

func (me *KeyInfoType) SetX509Data(value *X509Data) {
	me.X509Data = value
}

func (me *KeyInfoType) SetPGPData(value *PGPData) {
	me.PGPData = value
}

func (me *KeyInfoType) SetSPKIData(value *SPKIData) {
	me.SPKIData = value
}

func (me *KeyInfoType) SetMgmtData(value *MgmtData) {
	me.MgmtData = value
}
