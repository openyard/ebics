// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package xmldsig

// complex type
type DSAKeyValueType struct {
	G CryptoBinary `xml:"G,omitempty"`
	Y CryptoBinary `xml:"Y"`
	J CryptoBinary `xml:"J,omitempty"`
}
