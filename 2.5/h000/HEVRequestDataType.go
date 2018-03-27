// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
