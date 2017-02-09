package h004_test

import (
	"encoding/xml"
	"fmt"

	"github.com/openyard/ebics/2.5/h004"
)

func ExampleUnmarshalEbicsUnsignedRequest() {
	var req h004.EbicsUnsignedRequest
	xml.Unmarshal(unsignedRequest(), &req)
	fmt.Println(fmt.Sprintf("Version        :%10s", req.Version))
	fmt.Println(fmt.Sprintf("Revision       :%10d", req.Revision))
	fmt.Println(fmt.Sprintf("HostID         :%10s", req.Header.Static.HostID))
	fmt.Println(fmt.Sprintf("PartnerID      :%10s", req.Header.Static.PartnerID))
	fmt.Println(fmt.Sprintf("UserID         :%10s", req.Header.Static.UserID))
	fmt.Println(fmt.Sprintf("OrderType      :%10s", req.Header.Static.OrderDetails.OrderType))
	fmt.Println(fmt.Sprintf("OrderAttribute :%10s", req.Header.Static.OrderDetails.OrderAttribute))
	fmt.Println(fmt.Sprintf("SecurityMedium :%10s", req.Header.Static.SecurityMedium))
	fmt.Println(fmt.Sprintf("authenticate   :%10t", req.Body.DataTransfer.SignatureData.Authenticate))
	fmt.Println(fmt.Sprintf("SignatureData  :%10s", req.Body.DataTransfer.SignatureData.Value))
	fmt.Println(fmt.Sprintf("OrderData      :%10s", req.Body.DataTransfer.OrderData.Value))
	// Output:
	//Version        :      H004
	//Revision       :         1
	//HostID         :  EBIXHOST
	//PartnerID      :  KUNDE001
	//UserID         :    TLN100
	//OrderType      :       HSA
	//OrderAttribute :     OZNNN
	//SecurityMedium :      0000
	//authenticate   :      true
	//SignatureData  :      AQAB
	//OrderData      :      AQAB
}

func unsignedRequest() []byte {
	return []byte(`<?xml version="1.0" encoding="UTF-8"?>
<ebicsUnsignedRequest xmlns="urn:org:ebics:H004" xmlns:ds="http://www.w3.org/2000/09/xmldsig#" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="urn:org:ebics:H004 ebics_keymgmt_request_H004.xsd" Version="H004" Revision="1">
  <header authenticate="true">
    <static>
      <HostID>EBIXHOST</HostID>
      <PartnerID>KUNDE001</PartnerID>
      <UserID>TLN100</UserID>
      <OrderDetails>
        <OrderType>HSA</OrderType>
        <OrderAttribute>OZNNN</OrderAttribute>
      </OrderDetails>
      <SecurityMedium>0000</SecurityMedium>
    </static>
    <mutable/>
  </header>
  <body>
    <DataTransfer>
      <!--Zu ebics_orders_H004.xsd konformaes (XML-)Instanzdokument mit Wurzelelement
    UserSignatureData, kompimiert und base64-kodiert-->
      <SignatureData authenticate="true">AQAB</SignatureData>
      <!--Zu ebics_orders.xsd konformes (XML-)Instanzdokument mit Wurzelelement
    HSARequestOrderData, komprimiert und base64-kodiert-->
      <OrderData>AQAB</OrderData>
    </DataTransfer>
  </body>
</ebicsUnsignedRequest>`)
}
