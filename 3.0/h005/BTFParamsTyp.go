// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *BTFParamsTyp) SetSignatureFlag(value *SignatureFlagType) {
	me.SignatureFlag = value
}

func (me *BTFParamsTyp) SetDateRange(value *DateRangeType) {
	me.DateRange = value
}

func (me *BTFParamsTyp) SetParameter(value *Parameter) {
	me.Parameter = value
}
