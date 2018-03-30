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
										Element: NewElement("ID", "xsd:string", "unbounded", 0),
									},
									{
										Element: NewElement("RowID", "xsd:string", "unbounded", 0),
									},
									{
										Element: NewElement("CustNo", "xsd:string", "unbounded", 0),
									},
									{
										Element: NewElement("SubrNo", "xsd:string", "unbounded", 0),
									},
									{
										Element: NewElement("ListName", "xsd:string", "unbounded", 0),
									},
								},
							},
						},
					},
				},
			},
		},
		Messages: []wsdl.Message{
			NewMessage("ElasticInput", "Elastic", "GOWSDL_Message"),
			NewMessage("ElasticOutput", "ElasticResponse", "GOWSDL_Message"),
			NewMessage("ElasticError", "ElasticFault", "GOWSDL_Message"),
		},
		PortType: wsdl.PortType{
			Name: "GOWSDL_PortType",
			Operations: []wsdl.WSDLOperation{
				{
					Name:   "Elastic",
					Input:  NewIOOperation("ElasticInput", ""),
					Output: NewIOOperation("ElasticOutput", ""),
					Fault:  NewFaultOperation("ElasticError", "ElasticError", ""),
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
				NewWSDLOperation("Elastic", "Elastic", "ElasticError"),
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

func NewElement(name, typ, max string, min int) wsdl.Element {
	return wsdl.Element{
		Name:      name,
		Type:      typ,
		MinOccurs: min,
		MaxOccurs: max,
	}
}

func NewMessage(name, elem, msgName string) wsdl.Message {
	return wsdl.Message{
		Name: name,
		Part: wsdl.Part{
			Element: elem,
			Name:    msgName,
		},
	}
}

func NewWSDLOperation(name, action, fault string) wsdl.WSDLOperation {
	return wsdl.WSDLOperation{
		Name: name,
		Operation: &wsdl.ActionOperation{
			SoapAction: action,
			Style:      "document",
		},
		Input:  NewIOOperation("", "literal"),
		Output: NewIOOperation("", "literal"),
		Fault:  NewFaultOperation("", fault, "literal"),
	}
}

func NewIOOperation(msg, use string) wsdl.IOOperation {
	return wsdl.InputOperation{
		Operation: wsdl.Operation{
			Message: msg,
		},
		Body: NewSOAPBody(use),
	}
}

func NewSOAPBody(use string) *wsdl.SOAPBody {
	if use != "" {
		return &wsdl.SOAPBody{
			Use: use,
		}
	}
	return nil
}

func NewFaultOperation(msg, name, use string) wsdl.FaultOperation {
	return wsdl.FaultOperation{
		Operation: wsdl.Operation{
			Message: msg,
		},
		Name:  name,
		Fault: NewSOAPFault(name, use),
	}
}

func NewSOAPFault(name, use string) *wsdl.SOAPFault {
	if use != "" {
		return &wsdl.SOAPFault{
			Name: name,
			Use:  use,
		}
	}
	return nil
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
