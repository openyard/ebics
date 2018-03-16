// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

// ComplexType
type FULOrderParamsType struct {
	FileFormat *FileFormatType `xml:"FileFormat"`
	Parameter  *Parameter      `xml:"Parameter,omitempty"`
}

func NewFULOrderParamsType() *FULOrderParamsType {
	return new(FULOrderParamsType)
}

func (me *FULOrderParamsType) SetFileFormat(value *FileFormatType) {
	me.FileFormat = value
}

func (me *FULOrderParamsType) SetParameter(value *Parameter) {
	me.Parameter = value
}
