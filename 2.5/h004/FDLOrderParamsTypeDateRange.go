// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
