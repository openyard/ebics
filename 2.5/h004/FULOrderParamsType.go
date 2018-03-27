// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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

func (me *FULOrderParamsType) AddFileFormat() *FileFormatType {
	me.FileFormat = new(FileFormatType)
	return me.FileFormat
}

func (me *FULOrderParamsType) SetParameter(value *Parameter) {
	me.Parameter = value
}
