// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
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
