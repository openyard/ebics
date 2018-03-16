// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
