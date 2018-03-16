// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *EbicsKeyManagementResponseHeader) SetMutable(value *KeyMgmntResponseMutableHeaderType) {
	me.Mutable = value
}
