// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

// ComplexType
type FDLOrderParamsType struct {
	DateRange  *FDLOrderParamsTypeDateRange `xml:"DateRange,omitempty"`
	FileFormat *FileFormatType              `xml:"FileFormat"`
	Parameter  *Parameter                   `xml:"Parameter,omitempty"`
}

func NewFDLOrderParamsType() *FDLOrderParamsType {
	return new(FDLOrderParamsType)
}

func (me *FDLOrderParamsType) SetDateRange(value *FDLOrderParamsTypeDateRange) {
	me.DateRange = value
}

func (me *FDLOrderParamsType) AddDateRange() *FDLOrderParamsTypeDateRange {
	me.DateRange = new(FDLOrderParamsTypeDateRange)
	return me.DateRange
}

func (me *FDLOrderParamsType) SetFileFormat(value *FileFormatType) {
	me.FileFormat = value
}

func (me *FDLOrderParamsType) AddFileFormat() *FileFormatType {
	me.FileFormat = new(FileFormatType)
	return me.FileFormat
}

func (me *FDLOrderParamsType) SetParameter(value *Parameter) {
	me.Parameter = value
}
