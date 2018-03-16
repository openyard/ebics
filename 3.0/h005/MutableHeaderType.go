// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type MutableHeaderType struct {
	TransactionPhase *TransactionPhaseType           `xml:"TransactionPhase"`
	SegmentNumber    *MutableHeaderTypeSegmentNumber `xml:"SegmentNumber,omitempty"`

	Any []*w3c.Any
}

func NewMutableHeaderType() *MutableHeaderType {
	return new(MutableHeaderType)
}

func (me *MutableHeaderType) SetTransactionPhase(value *TransactionPhaseType) {
	me.TransactionPhase = value
}

func (me *MutableHeaderType) SetSegmentNumber(value *MutableHeaderTypeSegmentNumber) {
	me.SegmentNumber = value
}
