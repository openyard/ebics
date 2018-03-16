// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

// ComplexType
type EbicsNoPubKeyDigestsRequestHeader struct {
	Static  *NoPubKeyDigestsRequestStaticHeaderType `xml:"static"`
	Mutable *EmptyMutableHeaderType                 `xml:"mutable"`
	AuthenticationMarker
}

func NewEbicsNoPubKeyDigestsRequestHeader() *EbicsNoPubKeyDigestsRequestHeader {
	return new(EbicsNoPubKeyDigestsRequestHeader)
}

func (me *EbicsNoPubKeyDigestsRequestHeader) SetStatic(value *NoPubKeyDigestsRequestStaticHeaderType) {
	me.Static = value
}

func (me *EbicsNoPubKeyDigestsRequestHeader) SetMutable(value *EmptyMutableHeaderType) {
	me.Mutable = value
}
