// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

// ComplexType
type EbicsUnsecuredRequestHeader struct {
	Static  *UnsecuredRequestStaticHeaderType `xml:"static"`
	Mutable *EmptyMutableHeaderType           `xml:"mutable"`
	AuthenticationMarker
}

func NewEbicsUnsecuredRequestHeader() *EbicsUnsecuredRequestHeader {
	return new(EbicsUnsecuredRequestHeader)
}

func (me *EbicsUnsecuredRequestHeader) SetStatic(value *UnsecuredRequestStaticHeaderType) {
	me.Static = value
}

func (me *EbicsUnsecuredRequestHeader) SetMutable(value *EmptyMutableHeaderType) {
	me.Mutable = value
}
