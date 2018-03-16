// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type HVUSigningInfoType struct {
	ReadyToBeSigned *w3c.Boolean            `xml:"readyToBeSigned,attr"`
	NumSigRequired  *w3c.PositiveInteger    `xml:"NumSigRequired,attr"`
	NumSigDone      *w3c.NonNegativeInteger `xml:"NumSigDone,attr"`
}

func NewHVUSigningInfoType() *HVUSigningInfoType {
	return new(HVUSigningInfoType)
}

func (me *HVUSigningInfoType) SetReadyToBeSigned(value *w3c.Boolean) {
	me.ReadyToBeSigned = value
}

func (me *HVUSigningInfoType) SetNumSigRequired(value *w3c.PositiveInteger) {
	me.NumSigRequired = value
}

func (me *HVUSigningInfoType) SetNumSigDone(value *w3c.NonNegativeInteger) {
	me.NumSigDone = value
}
