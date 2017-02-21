// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
package xmldsig

// complex type
type DSAKeyValueType struct {
	G CryptoBinary `xml:"G,omitempty"`
	Y CryptoBinary `xml:"Y"`
	J CryptoBinary `xml:"J,omitempty"`
}
