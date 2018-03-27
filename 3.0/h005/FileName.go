// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

// Attribute
type FileName struct {
	Value *FileNameStringType `xml:"fileName,attr"`
}

func NewFileName() *FileName {
	return new(FileName)
}

func (me *FileName) SetValue(value *FileNameStringType) {
	me.Value = value
}
