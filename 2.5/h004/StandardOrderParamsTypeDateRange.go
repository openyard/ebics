// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *StandardOrderParamsTypeDateRange) SetEnd(value *DateType) {
	me.End = value
}
