// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package xmldsig

// complex type
type DSAKeyValueType struct {
	G CryptoBinary `xml:"G,omitempty"`
	Y CryptoBinary `xml:"Y"`
	J CryptoBinary `xml:"J,omitempty"`
}
