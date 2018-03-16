// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *HPDAccessParamsType) SetHostID(value *HostIDType) {
	me.HostID = value
}
