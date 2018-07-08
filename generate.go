package wsdl

import (
	"encoding/xml"

	"github.com/pallat/wsdl/def"
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
	Schema() def.Schema
	Messages() []def.Message
	PortTypeOperations() []def.WSDLOperation
	BindingOperation() def.WSDLOperation
	Service() def.Service
}

func NewOperation(pro Prototype) IOperation {
	return Operation{pro: pro}
}

type Operation struct {
	pro Prototype
}

func (o Operation) Schema() def.Schema {
	schema := def.DefaultSchema
	schema.Elements = o.elements()
	return schema
}

func (o Operation) elements() []def.SchemaElement {
	return []def.SchemaElement{
		{
			Element: def.Element{
				Name: o.pro.InputType().TypeName(),
			},
			ComplexType: def.ComplexType{
				Sequence: def.Sequence{
					Elements: NewElements(o.pro.InputType().SingleFields()),
				},
			},
		},
		{
			Element: def.Element{
				Name: o.pro.OutputType().TypeName(),
			},
			ComplexType: def.ComplexType{
				Sequence: def.Sequence{
					Elements: NewElements(o.pro.OutputType().SingleFields()),
				},
			},
		},
		{
			Element: def.Element{
				Name: o.pro.ErrorType().TypeName(),
			},
			ComplexType: def.ComplexType{
				Sequence: def.Sequence{
					Elements: NewElements(o.pro.ErrorType().SingleFields()),
				},
			},
		},
	}
}

func (o Operation) Messages() []def.Message {
	return []def.Message{
		def.NewMessage(o.pro.InputType().MessageName(), o.pro.InputType().TypeName()),
		def.NewMessage(o.pro.OutputType().MessageName(), o.pro.OutputType().TypeName()),
		def.NewMessage(o.pro.ErrorType().MessageName(), o.pro.ErrorType().TypeName()),
	}
}

func (o Operation) PortTypeOperations() []def.WSDLOperation {
	return []def.WSDLOperation{
		{
			Name:   o.pro.OperationName(),
			Input:  def.NewIOOperation(o.pro.InputType().MessageName(), ""),
			Output: def.NewIOOperation(o.pro.OutputType().MessageName(), ""),
			Fault:  def.NewFaultOperation(o.pro.ErrorType().MessageName(), o.pro.ErrorType().MessageName(), ""),
		},
	}
}

func (o Operation) BindingOperation() def.WSDLOperation {
	return def.NewWSDLOperation(o.pro.OperationName(), o.pro.OperationName(), o.pro.ErrorType().MessageName())
}

func (o Operation) Service() def.Service {
	return def.NewService(o.pro.Location())
}

func WSDL(opers ...IOperation) (string, error) {
	def := def.DefaultDefenitions

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

func NewElements(names []string) []def.Element {
	elements := []def.Element{}
	for _, v := range names {
		elements = append(elements, def.NewElement(v, "xsd:string", "0", "unbounded"))
	}
	return elements
}
