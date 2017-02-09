// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package xmldsig

// complex type
type DSAKeyValueType struct {
	G CryptoBinary `xml:"G,omitempty"`
	Y CryptoBinary `xml:"Y"`
	J CryptoBinary `xml:"J,omitempty"`
}
