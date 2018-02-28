// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type HPDProtocolParamsType struct {
	Version               HPDVersionType                             `xml:"Version"`
	Recovery              HPDProtocolParamsTypeRecovery              `xml:"Recovery,omitempty"`
	PreValidation         HPDProtocolParamsTypePreValidation         `xml:"PreValidation,omitempty"`
	ClientDataDownload    HPDProtocolParamsTypeClientDataDownload    `xml:"ClientDataDownload,omitempty"`
	DownloadableOrderData HPDProtocolParamsTypeDownloadableOrderData `xml:"DownloadableOrderData,omitempty"`

	Any []*w3c.Any
}
