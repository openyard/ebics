// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

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
