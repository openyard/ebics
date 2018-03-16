// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

// ComplexType
type EbicsRequestHeader struct {
	Static  *StaticHeaderType  `xml:"static"`
	Mutable *MutableHeaderType `xml:"mutable"`
	AuthenticationMarker
}

func NewEbicsRequestHeader() *EbicsRequestHeader {
	return new(EbicsRequestHeader)
}

func (me *EbicsRequestHeader) SetStatic(value *StaticHeaderType) {
	me.Static = value
}

func (me *EbicsRequestHeader) SetMutable(value *MutableHeaderType) {
	me.Mutable = value
}
