// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
