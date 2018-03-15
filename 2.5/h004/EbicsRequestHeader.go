// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h004

// complex type
type EbicsRequestHeader struct {
	Static  *StaticHeaderType  `xml:"static"`
	Mutable *MutableHeaderType `xml:"mutable"`
	AuthenticationMarker
}
