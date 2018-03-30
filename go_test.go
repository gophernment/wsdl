package wsdl

import (
	"fmt"
	"testing"
)

type Elastic struct{}

func (Elastic) Location() string {
	return "http://localhost:1323/elastic"
}
func (Elastic) OperationName() string {
	return "Elastic"
}

func (Elastic) InputType() Type {
	return ElasticInput{}
}
func (Elastic) OutputType() Type {
	return ElasticOutput{}
}
func (Elastic) ErrorType() Type {
	return ElasticError{}
}

type ElasticInput struct {
	ID       string `xml:"ID"`
	RowID    string `xml:"RowID"`
	CustNo   string `xml:"CustNo"`
	SubrNo   string `xml:"SubrNo"`
	ListName string `xml:"ListName"`
}

func (ElasticInput) Name() string {
	return "ElasticInput"
}
func (ElasticInput) SingleFields() []string {
	return []string{"ID", "RowID", "CustNo", "SubrNo", "ListName"}
}

type ElasticOutput struct{}

func (ElasticOutput) Name() string {
	return "ElasticOutput"
}
func (ElasticOutput) SingleFields() []string {
	return []string{"Index", "Type", "ID", "Version", "Created"}
}

type ElasticError struct{}

func (ElasticError) Name() string {
	return "ElasticError"
}
func (ElasticError) SingleFields() []string {
	return []string{"En", "Th", "Code"}
}

func TestGenEasyWSDLFromXMLString(t *testing.T) {
	wsdlString, err := WSDL(Elastic{})
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(wsdlString)

	if wsdlString == wsdlFromXMLString {
		t.Error("not")
	}
}

var xmlString = `<Envelope>
	<Header/>
	<Body>
		<Elastic>
			<ID>11115</ID>
			<RowID>111111</RowID>
			<CustNo>1</CustNo>
			<SubrNo>aaaaaaaaaa</SubrNo>
			<ListName>aaasd</ListName>
		</Elastic>
	</Body>
</Envelope>`

var wsdlFromXMLString = `<wsdl:definitions xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap11="http://schemas.xmlsoap.org/wsdl/soap/" xmlns:tuxtype="urn:pack.GOWSDL_typedef.golang" xmlns:tns="urn:GOWSDL.wsdl" xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/" name="GOWSDL" targetNamespace="urn:GOWSDL.wsdl">
<wsdl:documentation>Generated at 03-26-2018 15:44:07:838</wsdl:documentation>
<wsdl:types>
	<xsd:schema attributeFormDefault="unqualified" elementFormDefault="qualified" targetNamespace="urn:pack.GOWSDL_typedef.golang">
		<xsd:element name="Elastic">
			<xsd:complexType>
				<xsd:sequence>
					<xsd:element maxOccurs="unbounded" minOccurs="0" name="ID" type="xsd:string"></xsd:element>
					<xsd:element maxOccurs="unbounded" minOccurs="0" name="RowID" type="xsd:string"></xsd:element>
					<xsd:element maxOccurs="unbounded" minOccurs="0" name="CustNo" type="xsd:string"></xsd:element>
					<xsd:element maxOccurs="unbounded" minOccurs="0" name="SubrNo" type="xsd:string"></xsd:element>
					<xsd:element maxOccurs="unbounded" minOccurs="0" name="ListName" type="xsd:string"></xsd:element>
				</xsd:sequence>
			</xsd:complexType>
		</xsd:element>
		<xsd:element name="ElasticResponse">
			<xsd:complexType>
				<xsd:sequence>
					<xsd:element maxOccurs="unbounded" minOccurs="0" name="Index" type="xsd:string"></xsd:element>
					<xsd:element maxOccurs="unbounded" minOccurs="0" name="Type" type="xsd:string"></xsd:element>
					<xsd:element maxOccurs="unbounded" minOccurs="0" name="ID" type="xsd:string"></xsd:element>
					<xsd:element maxOccurs="unbounded" minOccurs="0" name="Version" type="xsd:string"></xsd:element>
					<xsd:element maxOccurs="unbounded" minOccurs="0" name="Created" type="xsd:string"></xsd:element>
				</xsd:sequence>
			</xsd:complexType>
		</xsd:element>
		<xsd:element name="ElasticFault">
			<xsd:complexType>
				<xsd:sequence>
					<xsd:element maxOccurs="1" minOccurs="0" name="Error" type="xsd:string"></xsd:element>
				</xsd:sequence>
			</xsd:complexType>
		</xsd:element>
	</xsd:schema>
</wsdl:types>
<wsdl:message name="ElasticInput">
	<wsdl:part element="tuxtype:Elastic" name="FML32"></wsdl:part>
</wsdl:message>
<wsdl:message name="ElasticOutput">
	<wsdl:part element="tuxtype:ElasticResponse" name="FML32"></wsdl:part>
</wsdl:message>
<wsdl:message name="ElasticError">
	<wsdl:part element="tuxtype:ElasticFault" name="FML32"></wsdl:part>
</wsdl:message>
<wsdl:portType name="GOWSDL_PortType">
	<wsdl:operation name="Elastic">
		<wsdl:input message="tns:ElasticInput"></wsdl:input>
		<wsdl:output message="tns:ElasticOutput"></wsdl:output>
		<wsdl:fault message="tns:ElasticError" name="ElasticError"></wsdl:fault>
	</wsdl:operation>
</wsdl:portType>
<wsdl:binding name="GOWSDL_Binding" type="tns:GOWSDL_PortType">
	<soap11:binding style="document" transport="http://schemas.xmlsoap.org/soap/http"></soap11:binding>
	<wsdl:operation name="Elastic">
		<soap11:operation soapAction="Elastic" style="document"></soap11:operation>
		<wsdl:input>
			<soap11:body use="literal"></soap11:body>
		</wsdl:input>
		<wsdl:output>
			<soap11:body use="literal"></soap11:body>
		</wsdl:output>
		<wsdl:fault name="ElasticError">
			<soap11:fault name="ElasticError" use="literal"></soap11:fault>
		</wsdl:fault>
	</wsdl:operation>
</wsdl:binding>
<wsdl:service name="TuxedoWebService">
	<wsdl:port binding="tns:GOWSDL_Binding" name="GOWSDL_Endpoint">
		<soap11:address location="http://localhost:1323/elastic"></soap11:address>
	</wsdl:port>
</wsdl:service>
</wsdl:definitions>`
