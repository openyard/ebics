// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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
