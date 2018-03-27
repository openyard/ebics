// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

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
