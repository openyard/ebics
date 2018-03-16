// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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
