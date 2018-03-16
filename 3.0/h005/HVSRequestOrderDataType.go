// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type HVSRequestOrderDataType struct {
	CancelledDataDigest *DataDigestType `xml:"CancelledDataDigest"`

	Any []*w3c.Any
}

func NewHVSRequestOrderDataType() *HVSRequestOrderDataType {
	return new(HVSRequestOrderDataType)
}

func (me *HVSRequestOrderDataType) SetCancelledDataDigest(value *DataDigestType) {
	me.CancelledDataDigest = value
}
