package def

func NewElement(name, typ, min, max string) Element {
	return Element{
		Name:      name,
		Type:      typ,
		MinOccurs: min,
		MaxOccurs: max,
	}
}

func NewMessage(name, elem string) Message {
	return Message{
		Name: name,
		Part: Part{
			Element: "gotype:" + elem,
			Name:    "GOWSDL_Message",
		},
	}
}

func NewWSDLOperation(name, action, fault string) WSDLOperation {
	return WSDLOperation{
		Name: name,
		Operation: &ActionOperation{
			SoapAction: action,
			Style:      "document",
		},
		Input:  NewIOOperation("", "literal"),
		Output: NewIOOperation("", "literal"),
		Fault:  NewFaultOperation("", fault, "literal"),
	}
}

func NewIOOperation(msg, use string) IOOperation {
	if msg != "" {
		msg = "tns:" + msg
	}
	return InputOperation{
		Operation: Operation{
			Message: msg,
		},
		Body: NewSOAPBody(use),
	}
}

func NewSOAPBody(use string) *SOAPBody {
	if use != "" {
		return &SOAPBody{
			Use: use,
		}
	}
	return nil
}

func NewFaultOperation(msg, name, use string) FaultOperation {
	if msg != "" {
		msg = "tns:" + msg
	}

	return FaultOperation{
		Operation: Operation{
			Message: msg,
		},
		Name:  name,
		Fault: NewSOAPFault(name, use),
	}
}

func NewSOAPFault(name, use string) *SOAPFault {
	if use != "" {
		return &SOAPFault{
			Name: name,
			Use:  use,
		}
	}
	return nil
}

func NewService(loc string) Service {
	return Service{
		Name: "GolangWebService",
		Port: Port{
			Binding: "tns:GOWSDL_Binding",
			Name:    "GOWSDL_Endpoint",
			Address: Address{
				Location: loc,
			},
		},
	}
}

var DefaultDefenitions = Definitions{
	WSDL:            "http://schemas.xmlsoap.org/wsdl/",
	XSD:             "http://www.w3.org/2001/XMLSchema",
	Soap:            "http://schemas.xmlsoap.org/wsdl/soap/",
	TNS:             "urn:GOWSDL.wsdl",
	TargetNamespace: "urn:GOWSDL.wsdl",
	TypeAttr:        "urn:pack.GOWSDL_typedef.golang",
	Documentation:   "",
	Types: Types{
		Schemas: []Schema{},
	},
	Messages: []Message{},
	PortType: DefaultPortType,
	Binding:  DefaultBinding,
	Service:  []Service{},
}

var DefaultPortType = PortType{
	Name:       "GOWSDL_PortType",
	Operations: []WSDLOperation{},
}

var DefaultSchema = Schema{
	AttributeFormDefault: "unqualified",
	ElementFormDefault:   "qualified",
	TargetNamespace:      "urn:pack.GOWSDL_typedef.golang",
	Imports:              nil,
	Elements:             []SchemaElement{},
}

var DefaultBinding = Binding{
	Name: "GOWSDL_Binding",
	Type: "tns:GOWSDL_PortType",
	Binding: SOAPBinding{
		Style:     "document",
		Transport: "http://schemas.xmlsoap.org/soap/http",
	},
	Operation: []WSDLOperation{},
}
