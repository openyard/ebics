// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

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

func (me *EbicsUnsecuredRequestHeader) AddStatic() *UnsecuredRequestStaticHeaderType {
	me.Static = new(UnsecuredRequestStaticHeaderType)
	return me.Static
}

func (me *EbicsUnsecuredRequestHeader) SetMutable(value *EmptyMutableHeaderType) {
	me.Mutable = value
}

func (me *EbicsUnsecuredRequestHeader) AddMutable() *EmptyMutableHeaderType {
	me.Mutable = new(EmptyMutableHeaderType)
	return me.Mutable
}
