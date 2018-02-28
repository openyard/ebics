// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type HPDProtocolParamsType struct {
	Version               HPDVersionType                             `xml:"Version"`
	Recovery              HPDProtocolParamsTypeRecovery              `xml:"Recovery,omitempty"`
	PreValidation         HPDProtocolParamsTypePreValidation         `xml:"PreValidation,omitempty"`
	X509Data              HPDProtocolParamsTypeX509Data              `xml:"X509Data,omitempty"`
	ClientDataDownload    HPDProtocolParamsTypeClientDataDownload    `xml:"ClientDataDownload,omitempty"`
	DownloadableOrderData HPDProtocolParamsTypeDownloadableOrderData `xml:"DownloadableOrderData,omitempty"`

	Any []*w3c.Any
}
