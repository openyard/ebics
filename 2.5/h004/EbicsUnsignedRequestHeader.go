// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

// ComplexType
type EbicsUnsignedRequestHeader struct {
	Static  *UnsignedRequestStaticHeaderType `xml:"static"`
	Mutable *EmptyMutableHeaderType          `xml:"mutable"`
	AuthenticationMarker
}

func NewEbicsUnsignedRequestHeader() *EbicsUnsignedRequestHeader {
	return new(EbicsUnsignedRequestHeader)
}

func (me *EbicsUnsignedRequestHeader) SetStatic(value *UnsignedRequestStaticHeaderType) {
	me.Static = value
}

func (me *EbicsUnsignedRequestHeader) SetMutable(value *EmptyMutableHeaderType) {
	me.Mutable = value
}
