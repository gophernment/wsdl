package wsdl

import (
	"encoding/xml"

	"github.com/pallat/wsdl"
)

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
	def := wsdl.DefaultDefenitions

	schemas := wsdl.DefaultSchema
	schemas.Elements = []wsdl.SchemaElement{
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
	}

	def.Types = wsdl.Types{
		Schemas: []wsdl.Schema{schemas},
	}

	def.Messages = []wsdl.Message{
		wsdl.NewMessage(pro.InputType().MessageName(), pro.InputType().TypeName()),
		wsdl.NewMessage(pro.OutputType().MessageName(), pro.OutputType().TypeName()),
		wsdl.NewMessage(pro.ErrorType().MessageName(), pro.ErrorType().TypeName()),
	}

	def.PortType.Operations = []wsdl.WSDLOperation{
		{
			Name:   pro.OperationName(),
			Input:  wsdl.NewIOOperation(pro.InputType().MessageName(), ""),
			Output: wsdl.NewIOOperation(pro.OutputType().MessageName(), ""),
			Fault:  wsdl.NewFaultOperation(pro.ErrorType().MessageName(), pro.ErrorType().MessageName(), ""),
		},
	}

	def.Binding.Operation = append(def.Binding.Operation, wsdl.NewWSDLOperation(pro.OperationName(), pro.OperationName(), pro.ErrorType().MessageName()))
	def.Service = wsdl.NewService(pro.Location())

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
