package wsdl

import (
	"encoding/xml"

	"github.com/pallat/wsdl"
)

type Prototype interface {
	Location() string      // url of service
	OperationName() string // operation name
	InputType() Type
	OutputType() Type
	ErrorType() Type
}

type Type interface {
	MessageName() string    // Alias name, request usually follow by Input, response usually follow by Output, error usually be follow by Error
	TypeName() string       // Type name of XML, request usually same with operation name, response usually follow by Response, error usually be follow by Fault
	SingleFields() []string // array of fields name
	// ArrayFields() []string
}

type IOperation interface {
	Schema() wsdl.Schema
	Messages() []wsdl.Message
	PortTypeOperations() []wsdl.WSDLOperation
	BindingOperation() wsdl.WSDLOperation
	Service() wsdl.Service
}

func NewOperation(pro Prototype) IOperation {
	return Operation{pro: pro}
}

type Operation struct {
	pro Prototype
}

func (o Operation) Schema() wsdl.Schema {
	schema := wsdl.DefaultSchema
	schema.Elements = o.elements()
	return schema
}

func (o Operation) elements() []wsdl.SchemaElement {
	return []wsdl.SchemaElement{
		{
			Element: wsdl.Element{
				Name: o.pro.InputType().TypeName(),
			},
			ComplexType: wsdl.ComplexType{
				Sequence: wsdl.Sequence{
					Elements: NewElements(o.pro.InputType().SingleFields()),
				},
			},
		},
		{
			Element: wsdl.Element{
				Name: o.pro.OutputType().TypeName(),
			},
			ComplexType: wsdl.ComplexType{
				Sequence: wsdl.Sequence{
					Elements: NewElements(o.pro.OutputType().SingleFields()),
				},
			},
		},
		{
			Element: wsdl.Element{
				Name: o.pro.ErrorType().TypeName(),
			},
			ComplexType: wsdl.ComplexType{
				Sequence: wsdl.Sequence{
					Elements: NewElements(o.pro.ErrorType().SingleFields()),
				},
			},
		},
	}
}

func (o Operation) Messages() []wsdl.Message {
	return []wsdl.Message{
		wsdl.NewMessage(o.pro.InputType().MessageName(), o.pro.InputType().TypeName()),
		wsdl.NewMessage(o.pro.OutputType().MessageName(), o.pro.OutputType().TypeName()),
		wsdl.NewMessage(o.pro.ErrorType().MessageName(), o.pro.ErrorType().TypeName()),
	}
}

func (o Operation) PortTypeOperations() []wsdl.WSDLOperation {
	return []wsdl.WSDLOperation{
		{
			Name:   o.pro.OperationName(),
			Input:  wsdl.NewIOOperation(o.pro.InputType().MessageName(), ""),
			Output: wsdl.NewIOOperation(o.pro.OutputType().MessageName(), ""),
			Fault:  wsdl.NewFaultOperation(o.pro.ErrorType().MessageName(), o.pro.ErrorType().MessageName(), ""),
		},
	}
}

func (o Operation) BindingOperation() wsdl.WSDLOperation {
	return wsdl.NewWSDLOperation(o.pro.OperationName(), o.pro.OperationName(), o.pro.ErrorType().MessageName())
}

func (o Operation) Service() wsdl.Service {
	return wsdl.NewService(o.pro.Location())
}

func WSDL(opers ...IOperation) (string, error) {
	def := wsdl.DefaultDefenitions

	for _, oper := range opers {
		def.Types.Schemas = append(def.Types.Schemas, oper.Schema())
		def.Messages = append(def.Messages, oper.Messages()...)
		def.PortType.Operations = append(def.PortType.Operations, oper.PortTypeOperations()...)
		def.Binding.Operation = append(def.Binding.Operation, oper.BindingOperation())
		def.Service = append(def.Service, oper.Service())
	}

	b, err := xml.MarshalIndent(&def, "", "	")
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
