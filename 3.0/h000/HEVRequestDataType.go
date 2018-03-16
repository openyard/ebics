// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h000

import w3c "github.com/openyard/ebics/3.0/w3c"

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
