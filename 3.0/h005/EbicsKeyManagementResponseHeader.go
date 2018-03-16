// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

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
