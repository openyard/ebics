// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
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
