package h004_test

import (
	"encoding/xml"
	"fmt"

	"github.com/openyard/ebics/2.5/h004"
)

func ExampleUnmarshalEbicsUnsecuredRequest() {
	var req h004.EbicsUnsecuredRequest
	xml.Unmarshal(unsecuredRequest(), &req)
	fmt.Println(fmt.Sprintf("Version        :%10s", req.Version))
	fmt.Println(fmt.Sprintf("Revision       :%10d", req.Revision))
	fmt.Println(fmt.Sprintf("HostID         :%10s", req.Header.Static.HostID))
	fmt.Println(fmt.Sprintf("PartnerID      :%10s", req.Header.Static.PartnerID))
	fmt.Println(fmt.Sprintf("UserID         :%10s", req.Header.Static.UserID))
	fmt.Println(fmt.Sprintf("OrderType      :%10s", req.Header.Static.OrderDetails.OrderType))
	fmt.Println(fmt.Sprintf("OrderAttribute :%10s", req.Header.Static.OrderDetails.OrderAttribute))
	fmt.Println(fmt.Sprintf("SecurityMedium :%10s", req.Header.Static.SecurityMedium))
	fmt.Println(fmt.Sprintf("OrderData      :%10s", req.Body.DataTransfer.OrderData.Value))
	// Output:
	//Version        :      H004
	//Revision       :         1
	//HostID         :  EBIXHOST
	//PartnerID      :  KUNDE001
	//UserID         :    TLN100
	//OrderType      :       INI
	//OrderAttribute :     DZNNN
	//SecurityMedium :      0200
	//OrderData      :      AQAB
}

func unsecuredRequest() []byte {
	return []byte(`<?xml version="1.0" encoding="UTF-8"?>
<ebicsUnsecuredRequest xmlns="urn:org:ebics:H004" xmlns:ds="http://www.w3.org/2000/09/xmldsig#" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="urn:org:ebics:H004 ebics_ keymgmt_request_H004.xsd" Version="H004" Revision="1">
  <header authenticate="true">
    <static>
      <HostID>EBIXHOST</HostID>
      <PartnerID>KUNDE001</PartnerID>
      <UserID>TLN100</UserID>
      <OrderDetails>
        <OrderType>INI</OrderType>
        <OrderAttribute>DZNNN</OrderAttribute>
      </OrderDetails>
      <SecurityMedium>0200</SecurityMedium>
    </static>
    <mutable/>
  </header>
  <body>
    <DataTransfer>
      <!--INI-Datei gemäß Kapitel 14.2.5.4, komprimiert und base64-kodiert-->
      <OrderData>AQAB</OrderData>
    </DataTransfer>
  </body>
</ebicsUnsecuredRequest>`)
}
