// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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

func (me *EbicsNoPubKeyDigestsRequestHeader) AddStatic() *NoPubKeyDigestsRequestStaticHeaderType {
	me.Static = new(NoPubKeyDigestsRequestStaticHeaderType)
	return me.Static
}

func (me *EbicsNoPubKeyDigestsRequestHeader) SetMutable(value *EmptyMutableHeaderType) {
	me.Mutable = value
}

func (me *EbicsNoPubKeyDigestsRequestHeader) AddMutable() *EmptyMutableHeaderType {
	me.Mutable = new(EmptyMutableHeaderType)
	return me.Mutable
}
