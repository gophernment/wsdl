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
	MessageName() string
	TypeName() string
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
								Name: pro.InputType().TypeName(),
							},
							ComplexType: wsdl.ComplexType{
								Sequence: wsdl.Sequence{
									Elements: NewElements(pro.InputType().SingleFields()),
								},
							},
						},
						{
							Element: wsdl.Element{
								Name: pro.OutputType().TypeName(),
							},
							ComplexType: wsdl.ComplexType{
								Sequence: wsdl.Sequence{
									Elements: NewElements(pro.OutputType().SingleFields()),
								},
							},
						},
						{
							Element: wsdl.Element{
								Name: pro.ErrorType().TypeName(),
							},
							ComplexType: wsdl.ComplexType{
								Sequence: wsdl.Sequence{
									Elements: NewElements(pro.ErrorType().SingleFields()),
								},
							},
						},
					},
				},
			},
		},
		Messages: []wsdl.Message{
			wsdl.NewMessage(pro.InputType().MessageName(), pro.InputType().TypeName()),
			wsdl.NewMessage(pro.OutputType().MessageName(), pro.OutputType().TypeName()),
			wsdl.NewMessage(pro.ErrorType().MessageName(), pro.ErrorType().TypeName()),
		},
		PortType: wsdl.PortType{
			Name: "GOWSDL_PortType",
			Operations: []wsdl.WSDLOperation{
				{
					Name:   pro.OperationName(),
					Input:  wsdl.NewIOOperation(pro.InputType().MessageName(), ""),
					Output: wsdl.NewIOOperation(pro.OutputType().MessageName(), ""),
					Fault:  wsdl.NewFaultOperation(pro.ErrorType().MessageName(), pro.ErrorType().MessageName(), ""),
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
				wsdl.NewWSDLOperation(pro.OperationName(), pro.OperationName(), pro.ErrorType().MessageName()),
			},
		},
		Service: wsdl.NewService(pro.Location()),
	}

	b, err := xml.MarshalIndent(&def, "", "    ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func NewElements(names []string) []wsdl.Element {
	elements := []wsdl.Element{}
	for _, v := range names {
		elements = append(elements, wsdl.NewElement(v, "xsd:string", "0", "unbounded"))
	}
	return elements
}
