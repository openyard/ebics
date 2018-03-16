// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type SPKIDataType struct {
	SPKISexp *w3c.Base64Binary `xml:"SPKISexp"`

	Any []*w3c.Any
}

func NewSPKIDataType() *SPKIDataType {
	return new(SPKIDataType)
}

func (me *SPKIDataType) SetSPKISexp(value *w3c.Base64Binary) {
	me.SPKISexp = value
}

func (me *SPKIDataType) AddSPKISexp() *w3c.Base64Binary {
	me.SPKISexp = new(w3c.Base64Binary)
	return me.SPKISexp
}
