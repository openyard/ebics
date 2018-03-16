// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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
