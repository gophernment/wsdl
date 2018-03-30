package wsdl

import (
	"encoding/xml"

	"github.com/pallat/wsdl"
)

func WSDL(s string) (string, error) {
	def := wsdl.Definitions{
		Documentation: "",
		Types: wsdl.Types{
			Schemas: []wsdl.Schema{
				{
					AttributeFormDefault: "unqualified",
					ElementFormDefault:   "qualified",
					TargetNamespace:      "",
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
			wsdl.NewMessage("ElasticInput", "Elastic", "GOWSDL_Message"),
			wsdl.NewMessage("ElasticOutput", "ElasticResponse", "GOWSDL_Message"),
			wsdl.NewMessage("ElasticError", "ElasticFault", "GOWSDL_Message"),
		},
		PortType: wsdl.PortType{
			Name: "GOWSDL_PortType",
			Operations: []wsdl.WSDLOperation{
				{
					Name:   "Elastic",
					Input:  wsdl.NewIOOperation("ElasticInput", ""),
					Output: wsdl.NewIOOperation("ElasticOutput", ""),
					Fault:  wsdl.NewFaultOperation("ElasticError", "ElasticError", ""),
				},
			},
		},
		Binding: wsdl.Binding{
			Name: "GOWSDL_Binding",
			Type: "GOWSDL_PortType",
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
