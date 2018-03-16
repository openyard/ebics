// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
