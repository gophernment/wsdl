package def

import (
	"encoding/xml"
)

type Definitions struct {
	XMLName         xml.Name  `xml:"wsdl:definitions"`
	WSDL            string    `xml:"xmlns:wsdl,attr"`
	XSD             string    `xml:"xmlns:xsd,attr"`
	Soap            string    `xml:"xmlns:soap11,attr"`
	TNS             string    `xml:"xmlns:tns,attr"`
	TypeAttr        string    `xml:"xmlns:gotype,attr"`
	TargetNamespace string    `xml:"targetNamespace,attr"`
	Documentation   string    `xml:"wsdl:documentation"`
	Types           Types     `xml:"wsdl:types"`
	Messages        []Message `xml:"wsdl:message"`
	PortType        PortType  `xml:"wsdl:portType"`
	Binding         Binding   `xml:"wsdl:binding"`
	Service         []Service `xml:"wsdl:service"`
}

type Types struct {
	Schemas []Schema `xml:"xsd:schema"`
}

type Message struct {
	Name string `xml:"name,attr"`
	Part Part   `xml:"wsdl:part"`
}

type PortType struct {
	Name       string          `xml:"name,attr"`
	Operations []WSDLOperation `xml:"wsdl:operation"`
}

type Binding struct {
	Name      string          `xml:"name,attr"`
	Type      string          `xml:"type,attr"`
	Binding   SOAPBinding     `xml:"soap11:binding"`
	Operation []WSDLOperation `xml:"wsdl:operation"`
}

type Service struct {
	Name string `xml:"name,attr"`
	Port Port   `xml:"wsdl:port"`
}

type Part struct {
	Element string `xml:"element,attr"`
	Name    string `xml:"name,attr"`
}

type SOAPBody struct {
	Use string `xml:"use,attr"`
}

type ActionOperation struct {
	SoapAction string `xml:"soapAction,attr"`
	Style      string `xml:"style,attr"`
}

type Operation struct {
	Message string `xml:"message,attr,omitempty"`
}

type IOOperation struct {
	Operation
	Body *SOAPBody `xml:"soap11:body,omitempty"`
}

type FaultOperation struct {
	Operation
	Name  string     `xml:"name,attr"`
	Fault *SOAPFault `xml:"soap11:fault,omitempty"`
}

type InputOperation = IOOperation
type OutputOperation = IOOperation

type SOAPFault struct {
	Name string `xml:"name,attr"`
	Use  string `xml:"use,attr"`
}

type WSDLOperation struct {
	Name      string           `xml:"name,attr"`
	Operation *ActionOperation `xml:"soap11:operation,omitempty"`
	Input     InputOperation   `xml:"wsdl:input"`
	Output    OutputOperation  `xml:"wsdl:output"`
	Fault     FaultOperation   `xml:"wsdl:fault"`
}

type Element struct {
	Name      string `xml:"name,attr"`
	Type      string `xml:"xsd:type,attr,omitempty"`
	MinOccurs string `xml:"xsd:minOccurs,attr,omitempty"`
	MaxOccurs string `xml:"xsd:maxOccurs,attr,omitempty"`
}

type SequenceElement = Element

type Sequence struct {
	Elements []SequenceElement `xml:"xsd:element"`
}

type ComplexType struct {
	Name     string   `xml:"xsd:name,attr,omitempty"`
	Sequence Sequence `xml:"xsd:sequence"`
}

type Import struct {
	Namespace string `xml:"xsd:namespace,attr"`
}

type SchemaElement struct {
	Element
	ComplexType ComplexType `xml:"xsd:complexType"`
}

type Schema struct {
	AttributeFormDefault string          `xml:"attributeFormDefault,attr"`
	ElementFormDefault   string          `xml:"elementFormDefault,attr"`
	TargetNamespace      string          `xml:"targetNamespace,attr"`
	Imports              []Import        `xml:"xsd:import"`
	Elements             []SchemaElement `xml:"xsd:element"`
	ComplexTypes         []ComplexType   `xml:"xsd:complexType"`
}

type Address struct {
	Location string `xml:"location,attr"`
}

type Port struct {
	Binding string  `xml:"binding,attr"`
	Name    string  `xml:"name,attr"`
	Address Address `xml:"soap11:address"`
}

type SOAPBinding struct {
	Style     string `xml:"style,attr"`
	Transport string `xml:"transport,attr"`
}
