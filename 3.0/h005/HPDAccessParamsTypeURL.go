// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
