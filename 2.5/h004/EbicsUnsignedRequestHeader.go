// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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

func (me *EbicsUnsignedRequestHeader) AddStatic() *UnsignedRequestStaticHeaderType {
	me.Static = new(UnsignedRequestStaticHeaderType)
	return me.Static
}

func (me *EbicsUnsignedRequestHeader) SetMutable(value *EmptyMutableHeaderType) {
	me.Mutable = value
}

func (me *EbicsUnsignedRequestHeader) AddMutable() *EmptyMutableHeaderType {
	me.Mutable = new(EmptyMutableHeaderType)
	return me.Mutable
}
