// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type HPDProtocolParamsType struct {
	Version               *HPDVersionType                             `xml:"Version"`
	Recovery              *HPDProtocolParamsTypeRecovery              `xml:"Recovery,omitempty"`
	PreValidation         *HPDProtocolParamsTypePreValidation         `xml:"PreValidation,omitempty"`
	X509Data              *HPDProtocolParamsTypeX509Data              `xml:"X509Data,omitempty"`
	ClientDataDownload    *HPDProtocolParamsTypeClientDataDownload    `xml:"ClientDataDownload,omitempty"`
	DownloadableOrderData *HPDProtocolParamsTypeDownloadableOrderData `xml:"DownloadableOrderData,omitempty"`

	Any []*w3c.Any
}
