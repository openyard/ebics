// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type HPDProtocolParamsType struct {
	Version               *HPDVersionType                             `xml:"Version"`
	Recovery              *HPDProtocolParamsTypeRecovery              `xml:"Recovery,omitempty"`
	PreValidation         *HPDProtocolParamsTypePreValidation         `xml:"PreValidation,omitempty"`
	X509Data              *HPDProtocolParamsTypeX509Data              `xml:"X509Data,omitempty"`
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

func (me *HPDProtocolParamsType) SetRecovery(value *HPDProtocolParamsTypeRecovery) {
	me.Recovery = value
}

func (me *HPDProtocolParamsType) SetPreValidation(value *HPDProtocolParamsTypePreValidation) {
	me.PreValidation = value
}

func (me *HPDProtocolParamsType) SetX509Data(value *HPDProtocolParamsTypeX509Data) {
	me.X509Data = value
}

func (me *HPDProtocolParamsType) SetClientDataDownload(value *HPDProtocolParamsTypeClientDataDownload) {
	me.ClientDataDownload = value
}

func (me *HPDProtocolParamsType) SetDownloadableOrderData(value *HPDProtocolParamsTypeDownloadableOrderData) {
	me.DownloadableOrderData = value
}
