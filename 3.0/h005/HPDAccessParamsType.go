// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type HPDAccessParamsType struct {
	URL       []*HPDAccessParamsTypeURL `xml:"URL"`
	Institute *w3c.NormalizedString     `xml:"Institute"`
	HostID    *HostIDType               `xml:"HostID,omitempty"`

	Any []*w3c.Any
}

func NewHPDAccessParamsType() *HPDAccessParamsType {
	return new(HPDAccessParamsType)
}

func (me *HPDAccessParamsType) AddURL(value *HPDAccessParamsTypeURL) {
	me.URL = append(me.URL, value)
}

func (me *HPDAccessParamsType) SetInstitute(value *w3c.NormalizedString) {
	me.Institute = value
}

func (me *HPDAccessParamsType) AddInstitute() *w3c.NormalizedString {
	me.Institute = new(w3c.NormalizedString)
	return me.Institute
}

func (me *HPDAccessParamsType) SetHostID(value *HostIDType) {
	me.HostID = value
}

func (me *HPDAccessParamsType) AddHostID() *HostIDType {
	me.HostID = new(HostIDType)
	return me.HostID
}
