// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package w3c

import "encoding/xml"

// Any Type
type Any struct {
	XMLName xml.Name
	Content string `xml:",innerxml,omitempty"`
}
