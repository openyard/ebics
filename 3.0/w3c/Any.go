// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
package w3c

import "encoding/xml"

// any type
type Any struct {
	XMLName xml.Name
	Content string `xml:",innerxml,omitempty"`
}
