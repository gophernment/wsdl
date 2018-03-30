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
						Operation: wsdl.Operation{
							Message: "ElasticInput",
						},
					},
					Output: wsdl.OutputOperation{
						Operation: wsdl.Operation{
							Message: "ElasticOutput",
						},
					},
					Fault: wsdl.FaultOperation{
						Operation: wsdl.Operation{
							Message: "ElasticError",
						},
						Name: "ElasticError",
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
					Operation: &wsdl.ActionOperation{
						SoapAction: "Elastic",
						Style:      "document",
					},
					Input:  NewIOOperation(),
					Output: NewIOOperation(),
					Fault:  NewFaultOperation("ElasticError"),
				},
			},
		},
		Service: NewService("http://localhost:1323/elastic"),
	}

	b, err := xml.Marshal(&def)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func NewIOOperation() wsdl.IOOperation {
	return wsdl.InputOperation{
		Operation: wsdl.Operation{
			Message: "",
		},
		Body: &wsdl.SOAPBody{
			Use: "literal",
		},
	}
}

func NewFaultOperation(name string) wsdl.FaultOperation {
	return wsdl.FaultOperation{
		Operation: wsdl.Operation{
			Message: "",
		},
		Name: name,
		Fault: &wsdl.SOAPFault{
			Name: name,
			Use:  "literal",
		},
	}
}

func NewService(loc string) wsdl.Service {
	return wsdl.Service{
		Name: "GolangWebService",
		Port: wsdl.Port{
			Binding: "GOWSDL_Binding",
			Name:    "GOWSDL_Endpoint",
			Address: wsdl.Address{
				Location: loc,
			},
		},
	}
}
