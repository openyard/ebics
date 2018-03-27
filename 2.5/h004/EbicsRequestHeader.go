// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

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

func (me *EbicsRequestHeader) AddStatic() *StaticHeaderType {
	me.Static = new(StaticHeaderType)
	return me.Static
}

func (me *EbicsRequestHeader) SetMutable(value *MutableHeaderType) {
	me.Mutable = value
}

func (me *EbicsRequestHeader) AddMutable() *MutableHeaderType {
	me.Mutable = new(MutableHeaderType)
	return me.Mutable
}
