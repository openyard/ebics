// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h000

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type HEVResponseDataType struct {
	SystemReturnCode *SystemReturnCodeType               `xml:"SystemReturnCode"`
	VersionNumber    []*HEVResponseDataTypeVersionNumber `xml:"VersionNumber,omitempty"`

	Any []*w3c.Any
}

func NewHEVResponseDataType() *HEVResponseDataType {
	return new(HEVResponseDataType)
}

func (me *HEVResponseDataType) SetSystemReturnCode(value *SystemReturnCodeType) {
	me.SystemReturnCode = value
}

func (me *HEVResponseDataType) AddVersionNumber(value *HEVResponseDataTypeVersionNumber) {
	me.VersionNumber = append(me.VersionNumber, value)
}
