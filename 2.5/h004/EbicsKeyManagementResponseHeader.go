// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// ComplexType
type EbicsKeyManagementResponseHeader struct {
	Static  *EbicsKeyManagementResponseHeaderStatic `xml:"static"`
	Mutable *KeyMgmntResponseMutableHeaderType      `xml:"mutable"`
	AuthenticationMarker
}

func NewEbicsKeyManagementResponseHeader() *EbicsKeyManagementResponseHeader {
	return new(EbicsKeyManagementResponseHeader)
}

func (me *EbicsKeyManagementResponseHeader) SetStatic(value *EbicsKeyManagementResponseHeaderStatic) {
	me.Static = value
}

func (me *EbicsKeyManagementResponseHeader) AddStatic() *EbicsKeyManagementResponseHeaderStatic {
	me.Static = new(EbicsKeyManagementResponseHeaderStatic)
	return me.Static
}

func (me *EbicsKeyManagementResponseHeader) SetMutable(value *KeyMgmntResponseMutableHeaderType) {
	me.Mutable = value
}

func (me *EbicsKeyManagementResponseHeader) AddMutable() *KeyMgmntResponseMutableHeaderType {
	me.Mutable = new(KeyMgmntResponseMutableHeaderType)
	return me.Mutable
}
