// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
