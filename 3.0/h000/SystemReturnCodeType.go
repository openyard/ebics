// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
