// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// ComplexType
type StandardOrderParamsType struct {
	DateRange *StandardOrderParamsTypeDateRange `xml:"DateRange,omitempty"`
}

func NewStandardOrderParamsType() *StandardOrderParamsType {
	return new(StandardOrderParamsType)
}

func (me *StandardOrderParamsType) SetDateRange(value *StandardOrderParamsTypeDateRange) {
	me.DateRange = value
}

func (me *StandardOrderParamsType) AddDateRange() *StandardOrderParamsTypeDateRange {
	me.DateRange = new(StandardOrderParamsTypeDateRange)
	return me.DateRange
}
