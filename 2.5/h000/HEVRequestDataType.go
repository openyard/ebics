// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h000

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type HEVRequestDataType struct {
	HostID *HostIDType `xml:"HostID"`

	Any []*w3c.Any
}

func NewHEVRequestDataType() *HEVRequestDataType {
	return new(HEVRequestDataType)
}

func (me *HEVRequestDataType) SetHostID(value *HostIDType) {
	me.HostID = value
}

func (me *HEVRequestDataType) AddHostID() *HostIDType {
	me.HostID = new(HostIDType)
	return me.HostID
}
