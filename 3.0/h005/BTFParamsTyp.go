// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

// ComplexType
type BTFParamsTyp struct {
	FileName      *FileNameStringType    `xml:"fileName,attr"`
	Service       *RestrictedServiceType `xml:"Service"`
	SignatureFlag *SignatureFlagType     `xml:"SignatureFlag,omitempty"`
	DateRange     *DateRangeType         `xml:"DateRange,omitempty"`
	Parameter     *Parameter             `xml:"Parameter,omitempty"`
}

func NewBTFParamsTyp() *BTFParamsTyp {
	return new(BTFParamsTyp)
}

func (me *BTFParamsTyp) SetFileName(value *FileNameStringType) {
	me.FileName = value
}

func (me *BTFParamsTyp) SetService(value *RestrictedServiceType) {
	me.Service = value
}

func (me *BTFParamsTyp) AddService() *RestrictedServiceType {
	me.Service = new(RestrictedServiceType)
	return me.Service
}

func (me *BTFParamsTyp) SetSignatureFlag(value *SignatureFlagType) {
	me.SignatureFlag = value
}

func (me *BTFParamsTyp) AddSignatureFlag() *SignatureFlagType {
	me.SignatureFlag = new(SignatureFlagType)
	return me.SignatureFlag
}

func (me *BTFParamsTyp) SetDateRange(value *DateRangeType) {
	me.DateRange = value
}

func (me *BTFParamsTyp) AddDateRange() *DateRangeType {
	me.DateRange = new(DateRangeType)
	return me.DateRange
}

func (me *BTFParamsTyp) SetParameter(value *Parameter) {
	me.Parameter = value
}
