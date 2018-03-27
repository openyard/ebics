// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type HPDProtocolParamsType struct {
	Version               *HPDVersionType                             `xml:"Version"`
	Recovery              *HPDProtocolParamsTypeRecovery              `xml:"Recovery,omitempty"`
	PreValidation         *HPDProtocolParamsTypePreValidation         `xml:"PreValidation,omitempty"`
	ClientDataDownload    *HPDProtocolParamsTypeClientDataDownload    `xml:"ClientDataDownload,omitempty"`
	DownloadableOrderData *HPDProtocolParamsTypeDownloadableOrderData `xml:"DownloadableOrderData,omitempty"`

	Any []*w3c.Any
}

func NewHPDProtocolParamsType() *HPDProtocolParamsType {
	return new(HPDProtocolParamsType)
}

func (me *HPDProtocolParamsType) SetVersion(value *HPDVersionType) {
	me.Version = value
}

func (me *HPDProtocolParamsType) AddVersion() *HPDVersionType {
	me.Version = new(HPDVersionType)
	return me.Version
}

func (me *HPDProtocolParamsType) SetRecovery(value *HPDProtocolParamsTypeRecovery) {
	me.Recovery = value
}

func (me *HPDProtocolParamsType) AddRecovery() *HPDProtocolParamsTypeRecovery {
	me.Recovery = new(HPDProtocolParamsTypeRecovery)
	return me.Recovery
}

func (me *HPDProtocolParamsType) SetPreValidation(value *HPDProtocolParamsTypePreValidation) {
	me.PreValidation = value
}

func (me *HPDProtocolParamsType) AddPreValidation() *HPDProtocolParamsTypePreValidation {
	me.PreValidation = new(HPDProtocolParamsTypePreValidation)
	return me.PreValidation
}

func (me *HPDProtocolParamsType) SetClientDataDownload(value *HPDProtocolParamsTypeClientDataDownload) {
	me.ClientDataDownload = value
}

func (me *HPDProtocolParamsType) AddClientDataDownload() *HPDProtocolParamsTypeClientDataDownload {
	me.ClientDataDownload = new(HPDProtocolParamsTypeClientDataDownload)
	return me.ClientDataDownload
}

func (me *HPDProtocolParamsType) SetDownloadableOrderData(value *HPDProtocolParamsTypeDownloadableOrderData) {
	me.DownloadableOrderData = value
}

func (me *HPDProtocolParamsType) AddDownloadableOrderData() *HPDProtocolParamsTypeDownloadableOrderData {
	me.DownloadableOrderData = new(HPDProtocolParamsTypeDownloadableOrderData)
	return me.DownloadableOrderData
}
