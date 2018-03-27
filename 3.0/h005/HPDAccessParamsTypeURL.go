// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexElement
type HPDAccessParamsTypeURL struct {
	Value      *w3c.AnyURI    `xml:",chardata"`
	Valid_from *TimestampType `xml:"Valid_from,attr"`
}

func NewHPDAccessParamsTypeURL() *HPDAccessParamsTypeURL {
	return new(HPDAccessParamsTypeURL)
}

func (me *HPDAccessParamsTypeURL) SetValid_from(value *TimestampType) {
	me.Valid_from = value
}
