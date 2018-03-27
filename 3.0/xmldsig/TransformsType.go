// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
