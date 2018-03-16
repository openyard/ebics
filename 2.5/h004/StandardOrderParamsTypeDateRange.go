// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

// ComplexType
type StandardOrderParamsTypeDateRange struct {
	Start *DateType `xml:"Start"`
	End   *DateType `xml:"End"`
}

func NewStandardOrderParamsTypeDateRange() *StandardOrderParamsTypeDateRange {
	return new(StandardOrderParamsTypeDateRange)
}

func (me *StandardOrderParamsTypeDateRange) SetStart(value *DateType) {
	me.Start = value
}

func (me *StandardOrderParamsTypeDateRange) AddStart() *DateType {
	me.Start = new(DateType)
	return me.Start
}

func (me *StandardOrderParamsTypeDateRange) SetEnd(value *DateType) {
	me.End = value
}

func (me *StandardOrderParamsTypeDateRange) AddEnd() *DateType {
	me.End = new(DateType)
	return me.End
}
