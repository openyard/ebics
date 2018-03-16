// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *FDLOrderParamsTypeDateRange) SetEnd(value *DateType) {
	me.End = value
}
