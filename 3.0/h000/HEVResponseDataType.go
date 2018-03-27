// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h000

import w3c "github.com/openyard/ebics/3.0/w3c"

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

func (me *HEVResponseDataType) AddSystemReturnCode() *SystemReturnCodeType {
	me.SystemReturnCode = new(SystemReturnCodeType)
	return me.SystemReturnCode
}

func (me *HEVResponseDataType) AddVersionNumber(value *HEVResponseDataTypeVersionNumber) {
	me.VersionNumber = append(me.VersionNumber, value)
}
