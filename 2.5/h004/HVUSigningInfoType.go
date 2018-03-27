// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

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
