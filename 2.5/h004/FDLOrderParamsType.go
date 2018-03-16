// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *FDLOrderParamsType) SetFileFormat(value *FileFormatType) {
	me.FileFormat = value
}

func (me *FDLOrderParamsType) SetParameter(value *Parameter) {
	me.Parameter = value
}
