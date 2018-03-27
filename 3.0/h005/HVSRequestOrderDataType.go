// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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

func (me *HVSRequestOrderDataType) AddCancelledDataDigest() *DataDigestType {
	me.CancelledDataDigest = new(DataDigestType)
	return me.CancelledDataDigest
}
