// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package xmldsig

// complex type
type DSAKeyValueType struct {
	G *CryptoBinary `xml:"G,omitempty"`
	Y *CryptoBinary `xml:"Y"`
	J *CryptoBinary `xml:"J,omitempty"`
}
