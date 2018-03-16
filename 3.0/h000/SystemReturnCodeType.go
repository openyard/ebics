// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h000

// ComplexType
type SystemReturnCodeType struct {
	ReturnCode *ReturnCodeType `xml:"ReturnCode"`
	ReportText *ReportTextType `xml:"ReportText"`
}

func NewSystemReturnCodeType() *SystemReturnCodeType {
	return new(SystemReturnCodeType)
}

func (me *SystemReturnCodeType) SetReturnCode(value *ReturnCodeType) {
	me.ReturnCode = value
}

func (me *SystemReturnCodeType) SetReportText(value *ReportTextType) {
	me.ReportText = value
}
