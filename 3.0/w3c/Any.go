// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package w3c

import "encoding/xml"

// any type
type Any struct {
	XMLName xml.Name
	Content string `xml:",innerxml,omitempty"`
}
