// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
