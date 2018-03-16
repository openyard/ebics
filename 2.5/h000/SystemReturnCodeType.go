// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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

func (me *SystemReturnCodeType) AddReturnCode() *ReturnCodeType {
	me.ReturnCode = new(ReturnCodeType)
	return me.ReturnCode
}

func (me *SystemReturnCodeType) SetReportText(value *ReportTextType) {
	me.ReportText = value
}

func (me *SystemReturnCodeType) AddReportText() *ReportTextType {
	me.ReportText = new(ReportTextType)
	return me.ReportText
}
