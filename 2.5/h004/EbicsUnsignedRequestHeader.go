// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
