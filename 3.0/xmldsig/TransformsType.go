// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package xmldsig

// ComplexType
type TransformsType struct {
	Transform *Transform `xml:"Transform"`
}

func NewTransformsType() *TransformsType {
	return new(TransformsType)
}

func (me *TransformsType) SetTransform(value *Transform) {
	me.Transform = value
}
