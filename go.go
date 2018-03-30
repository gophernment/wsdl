package wsdl

import (
	"encoding/xml"

	"github.com/pallat/wsdl"
)

var serviceLocation = "http://localhost:1323/elastic"

type Prototype interface {
	Location() string
	OperationName() string
	InputType() Type
	OutputType() Type
	ErrorType() Type
}

type Type interface {
	Name() string
	SingleFields() []string
	// ArrayFields() []string
}

func WSDL(pro Prototype) (string, error) {
	def := wsdl.Definitions{
		WSDL:            "http://schemas.xmlsoap.org/wsdl/",
		XSD:             "http://www.w3.org/2001/XMLSchema",
		Soap:            "http://schemas.xmlsoap.org/wsdl/soap/",
		TNS:             "urn:GOWSDL.wsdl",
		TargetNamespace: "urn:GOWSDL.wsdl",
		TypeAttr:        "urn:pack.GOWSDL_typedef.golang",
		Documentation:   "",
		Types: wsdl.Types{
			Schemas: []wsdl.Schema{
				{
					AttributeFormDefault: "unqualified",
					ElementFormDefault:   "qualified",
					TargetNamespace:      "urn:pack.GOWSDL_typedef.golang",
					Imports:              nil,
					Elements: []wsdl.SchemaElement{
						{
							Element: wsdl.Element{
								Name: "Elastic",
							},
							ComplexType: wsdl.ComplexType{
								Name: "",
								Sequence: wsdl.Sequence{
									Elements: []wsdl.Element{
										wsdl.NewElement("ID", "xsd:string", "0", "unbounded"),
										wsdl.NewElement("RowID", "xsd:string", "0", "unbounded"),
										wsdl.NewElement("CustNo", "xsd:string", "0", "unbounded"),
										wsdl.NewElement("SubrNo", "xsd:string", "0", "unbounded"),
										wsdl.NewElement("ListName", "xsd:string", "0", "unbounded"),
									},
								},
							},
						},
					},
				},
			},
		},
		Messages: []wsdl.Message{
			wsdl.NewMessage("ElasticInput", "gotype:Elastic", "GOWSDL_Message"),
			wsdl.NewMessage("ElasticOutput", "gotype:ElasticResponse", "GOWSDL_Message"),
			wsdl.NewMessage("ElasticError", "gotype:ElasticFault", "GOWSDL_Message"),
		},
		PortType: wsdl.PortType{
			Name: "GOWSDL_PortType",
			Operations: []wsdl.WSDLOperation{
				{
					Name:   "Elastic",
					Input:  wsdl.NewIOOperation("tns:ElasticInput", ""),
					Output: wsdl.NewIOOperation("tns:ElasticOutput", ""),
					Fault:  wsdl.NewFaultOperation("tns:ElasticError", "ElasticError", ""),
				},
			},
		},
		Binding: wsdl.Binding{
			Name: "GOWSDL_Binding",
			Type: "tns:GOWSDL_PortType",
			Binding: wsdl.SOAPBinding{
				Style:     "document",
				Transport: "http://schemas.xmlsoap.org/soap/http",
			},
			Operation: []wsdl.WSDLOperation{
				wsdl.NewWSDLOperation("Elastic", "Elastic", "ElasticError"),
			},
		},
		Service: wsdl.NewService("http://localhost:1323/elastic"),
	}

	b, err := xml.MarshalIndent(&def, "", "    ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
