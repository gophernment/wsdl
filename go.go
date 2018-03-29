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
								Name:      "Elastic",
								Type:      "",
								MinOccurs: 0,
								MaxOccurs: "",
							},
						},
					},
					ComplexTypes: []wsdl.ComplexType{
						{
							Name: "",
							Sequence: wsdl.Sequence{
								Elements: []wsdl.SequenceElement{
									{
										Element: wsdl.Element{
											Name:      "ID",
											Type:      "xsd:string",
											MinOccurs: 0,
											MaxOccurs: "unbounded",
										},
									},
									{
										Element: wsdl.Element{
											Name:      "RowID",
											Type:      "xsd:string",
											MinOccurs: 0,
											MaxOccurs: "unbounded",
										},
									},
									{
										Element: wsdl.Element{
											Name:      "CustNo",
											Type:      "xsd:string",
											MinOccurs: 0,
											MaxOccurs: "unbounded",
										},
									},
									{
										Element: wsdl.Element{
											Name:      "SubrNo",
											Type:      "xsd:string",
											MinOccurs: 0,
											MaxOccurs: "unbounded",
										},
									},
									{
										Element: wsdl.Element{
											Name:      "ListName",
											Type:      "xsd:string",
											MinOccurs: 0,
											MaxOccurs: "unbounded",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		Messages: []wsdl.Message{
			{
				Name: "ElasticInput",
				Part: wsdl.Part{
					Element: "Elastic",
					Name:    "GOWSDL_Message",
				},
			},
		},
		PortType: wsdl.PortType{
			Name: "GOWSDL_PortType",
			Operations: []wsdl.WSDLOperation{
				{
					Name: "Elastic",
					Input: wsdl.InputOperation{
						IOOperation: wsdl.IOOperation{
							Operation: wsdl.Operation{
								Message: "ElasticInput",
							},
						},
					},
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
				{
					Name: "Elastic",
					Operation: wsdl.ActionOperation{
						SoapAction: "Elastic",
						Style:      "document",
					},
					Input: wsdl.InputOperation{
						IOOperation: wsdl.IOOperation{
							Operation: wsdl.Operation{
								Message: "",
							},
							Body: &wsdl.SOAPBody{
								Use: "literal",
							},
						},
					},
					// Output OutputOperation
					// Fault  FaultOperation
				},
			},
		},
		Service: wsdl.Service{
			Name: "GolangWebService",
			Port: wsdl.Port{
				Binding: "GOWSDL_Binding",
				Name:    "GOWSDL_Endpoint",
				Address: wsdl.Address{
					Location: "http://localhost:1323/elastic",
				},
			},
		},
	}

	b, err := xml.Marshal(&def)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
