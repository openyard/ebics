// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

// ComplexType
type DateRangeType struct {
	Start *DateType `xml:"Start"`
	End   *DateType `xml:"End"`
}

func NewDateRangeType() *DateRangeType {
	return new(DateRangeType)
}

func (me *DateRangeType) SetStart(value *DateType) {
	me.Start = value
}

func (me *DateRangeType) AddStart() *DateType {
	me.Start = new(DateType)
	return me.Start
}

func (me *DateRangeType) SetEnd(value *DateType) {
	me.End = value
}

func (me *DateRangeType) AddEnd() *DateType {
	me.End = new(DateType)
	return me.End
}
