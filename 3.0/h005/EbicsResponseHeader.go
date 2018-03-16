// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

// ComplexType
type EbicsResponseHeader struct {
	Static  *ResponseStaticHeaderType  `xml:"static"`
	Mutable *ResponseMutableHeaderType `xml:"mutable"`
	AuthenticationMarker
}

func NewEbicsResponseHeader() *EbicsResponseHeader {
	return new(EbicsResponseHeader)
}

func (me *EbicsResponseHeader) SetStatic(value *ResponseStaticHeaderType) {
	me.Static = value
}

func (me *EbicsResponseHeader) SetMutable(value *ResponseMutableHeaderType) {
	me.Mutable = value
}
