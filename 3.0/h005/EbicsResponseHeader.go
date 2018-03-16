// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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

func (me *EbicsResponseHeader) AddStatic() *ResponseStaticHeaderType {
	me.Static = new(ResponseStaticHeaderType)
	return me.Static
}

func (me *EbicsResponseHeader) SetMutable(value *ResponseMutableHeaderType) {
	me.Mutable = value
}

func (me *EbicsResponseHeader) AddMutable() *ResponseMutableHeaderType {
	me.Mutable = new(ResponseMutableHeaderType)
	return me.Mutable
}
