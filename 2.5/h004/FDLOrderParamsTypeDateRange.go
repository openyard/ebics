// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// ComplexType
type FDLOrderParamsTypeDateRange struct {
	Start *DateType `xml:"Start"`
	End   *DateType `xml:"End"`
}

func NewFDLOrderParamsTypeDateRange() *FDLOrderParamsTypeDateRange {
	return new(FDLOrderParamsTypeDateRange)
}

func (me *FDLOrderParamsTypeDateRange) SetStart(value *DateType) {
	me.Start = value
}

func (me *FDLOrderParamsTypeDateRange) AddStart() *DateType {
	me.Start = new(DateType)
	return me.Start
}

func (me *FDLOrderParamsTypeDateRange) SetEnd(value *DateType) {
	me.End = value
}

func (me *FDLOrderParamsTypeDateRange) AddEnd() *DateType {
	me.End = new(DateType)
	return me.End
}
