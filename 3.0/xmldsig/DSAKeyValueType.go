// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
package xmldsig

// complex type
type DSAKeyValueType struct {
	G CryptoBinary `xml:"G,omitempty"`
	Y CryptoBinary `xml:"Y"`
	J CryptoBinary `xml:"J,omitempty"`
}
