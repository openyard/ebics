// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package w3c

import "encoding/xml"

// any type
type Any struct {
	XMLName xml.Name
	Content string `xml:",innerxml,omitempty"`
}
